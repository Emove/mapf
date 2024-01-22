package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"reflect"
	"sync"
	"time"
)

const AmqpScheme = "amqp"

type rabbitConfiguration struct {
	addr   string
	mutex  sync.Mutex
	conn   *amqp.Connection
	logger log.Logger
}

func (config *rabbitConfiguration) GetAddr() string {
	return config.addr
}

// NewRabbitMqConfiguration 创建RabbitMQ配置对象
func NewRabbitMqConfiguration(addr string, logger log.Logger) (Configuration, func(), error) {
	config := &rabbitConfiguration{
		addr:   addr,
		logger: logger,
	}
	conn, err := amqp.Dial(config.addr)
	if err != nil {
		return nil, func() {}, err
	}
	config.mutex = sync.Mutex{}
	config.conn = conn
	return config, func() {
		_ = conn.Close()
	}, nil
}

var _ Publisher = (*rabbitMqChannel)(nil)
var _ Subscriber = (*rabbitMqChannel)(nil)

type rabbitMqChannel struct {
	config            *rabbitConfiguration
	mutex             sync.Mutex
	closed            bool
	channel           *amqp.Channel
	exchange          string
	key               string
	queue             string
	logger            *log.Helper
	channelClosedChan chan *amqp.Error
	done              chan struct{}
	deliveries        <-chan amqp.Delivery
	consumer          Consumer
	reconsume         chan struct{}
}

func NewRabbitMQPublisher(config Configuration, exchange, key, queue string) (Publisher, func(), error) {
	return newRabbitMqChannel(config, exchange, key, queue)
}

func NewRabbitMqSubscriber(config Configuration, exchange, key, queue string) (Subscriber, func(), error) {
	return newRabbitMqChannel(config, exchange, key, queue)
}

func newRabbitMqChannel(config Configuration, exchange, key, queue string) (*rabbitMqChannel, func(), error) {
	conf, ok := config.(*rabbitConfiguration)
	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("RabbitMq configuration excepted, got %v", reflect.TypeOf(config)))
	}
	conf.mutex.Lock()
	ch, err := conf.conn.Channel()
	if err != nil {
		return nil, func() {}, err
	}
	conf.mutex.Unlock()
	helper := log.NewHelper(conf.logger)
	channel := &rabbitMqChannel{
		config:            conf,
		mutex:             sync.Mutex{},
		channel:           ch,
		exchange:          exchange,
		key:               key,
		queue:             queue,
		logger:            helper,
		channelClosedChan: make(chan *amqp.Error),
		done:              make(chan struct{}),
	}
	err = declareDirectExchange(ch, helper, exchange)
	if err != nil {
		return nil, channel.close, err
	}
	err = declareAndBindQueue(ch, helper, exchange, key, queue)
	if err != nil {
		return nil, channel.close, err
	}
	if err = ch.Confirm(false); err != nil {
		return nil, channel.close, err
	}
	//channel.confirmCh = ch.NotifyPublish(make(chan amqp.Confirmation))
	//ch.NotifyReturn()
	ch.NotifyClose(channel.channelClosedChan)
	go channel.keepalive()
	return channel, channel.close, nil
}

func (ch *rabbitMqChannel) Publish(ctx context.Context, msg interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err = ch.ensureChannelOpened(); err != nil {
		return err
	}

	err = ch.channel.PublishWithContext(
		ctx,
		ch.exchange,
		ch.key,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		})
	if err != nil {
		return err
	}
	ch.logger.Infof("publish event to queue: %s, event: %s", ch.queue, string(body))
	return nil
}

func (ch *rabbitMqChannel) Subscribe(consumer Consumer) error {
	deliveries, err := ch.channel.ConsumeWithContext(context.Background(), ch.queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	ch.deliveries = deliveries
	ch.consumer = consumer
	ch.reconsume = make(chan struct{})
	go func() {
		for {
			select {
			case <-ch.done:
				close(ch.reconsume)
				return
			case <-ch.reconsume:
				ch.deliveries, _ = ch.channel.ConsumeWithContext(context.Background(), ch.queue, "", false, false, false, false, nil)
			case delivery := <-ch.deliveries:
				ch.consume(delivery)
			}
		}
	}()
	return nil
}

func (ch *rabbitMqChannel) consume(delivery amqp.Delivery) {
	defer func() {
		if err := recover(); err != nil {
			_ = delivery.Nack(false, true)
		}
	}()
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	ch.logger.Infof("consume event from queue: %s, event: %s", ch.queue, string(delivery.Body))
	err := ch.consumer(timeout, delivery.Body)
	if err != nil {
		ch.logger.Errorw("subscriber", ch.queue, "deliveryTag", delivery.DeliveryTag, "msg", delivery.Body, "err", err)
		_ = delivery.Nack(false, true)
	}
	_ = delivery.Ack(false)
}

func (ch *rabbitMqChannel) close() {
	ch.mutex.Lock()
	defer ch.mutex.Unlock()
	ch.closed = true
	_ = ch.channel.Close()
	close(ch.done)
}

func (ch *rabbitMqChannel) keepalive() {
	for {
		select {
		case <-ch.done:
			return
		case err := <-ch.channelClosedChan:
			if err == nil && !ch.channel.IsClosed() {
				continue
			}
			ch.mutex.Lock()
			if ch.closed == true {
				ch.mutex.Unlock()
				continue
			}
			ch.mutex.Unlock()
			ch.logger.Errorf("channel closed: %v", err)
			connectErr := ch.ensureChannelOpened()
			if connectErr != nil {
				ch.logger.Errorf("reconnect error: %v", connectErr)
				time.Sleep(5 * time.Second)
				continue
			}
			if ch.consumer != nil {
				ch.reconsume <- struct{}{}
			}
		}
	}
}

func (ch *rabbitMqChannel) ensureChannelOpened() (err error) {
	if !ch.channel.IsClosed() {
		return
	}
	ch.mutex.Lock()
	defer ch.mutex.Unlock()
	if !ch.channel.IsClosed() {
		return nil
	}
	logger := ch.logger
	config := ch.config
	config.mutex.Lock()
	defer config.mutex.Unlock()
	if config.conn.IsClosed() {
		if config.conn, err = amqp.Dial(config.addr); err != nil {
			return err
		}
		logger.Infof("reconnected to RabbitMQ, url: %s", config.addr)
	}
	if ch.channel, err = ch.config.conn.Channel(); err != nil {
		return err
	}
	logger.Infof("reopen channel for queue: %s", ch.queue)
	return nil
}

func declareDirectExchange(ch *amqp.Channel, logger *log.Helper, exchange string) error {
	if err := ch.ExchangeDeclare(exchange, amqp.ExchangeDirect, true, false, false, false, nil); err != nil {
		return err
	}
	logger.Infof("declared exchange, name: %s", exchange)
	return nil
}

func declareAndBindQueue(ch *amqp.Channel, logger *log.Helper, exchange, key, queue string) error {
	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	logger.Infof("declared queue, name: %s, messages: %d, consumers: %d", q.Name, q.Messages, q.Consumers)
	err = ch.QueueBind(queue, key, exchange, false, nil)
	if err == nil {
		logger.Infof("bind queue to exchange, exchange: %s, key: %s, queue: %s", exchange, key, queue)
	}
	return err
}
