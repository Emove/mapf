package data

import (
	plan "mapf/api/plan/v1"
	robot "mapf/api/robot/v1"
	warehouse "mapf/api/warehouse/v1"
	"mapf/app/web/internal/conf"
	"mapf/internal/data"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	Whc warehouse.WarehouseServiceClient
	Rbc robot.RobotServiceClient
	Pac plan.PlanServiceClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)
	instance := &Data{}
	var err error
	for _, service := range c.GetServices() {
		serviceName, addr, timeout := strings.ToLower(service.Name), service.Addr, service.Timeout.AsDuration()
		if serviceName == "warehouse" {
			instance.Whc, err = newWarehouseClient(addr, timeout)
		} else if serviceName == "robot" {
			instance.Rbc, err = newRobotClient(addr, timeout)
		} else if serviceName == "plan" {
			instance.Pac, err = newPlanClient(addr, timeout)
		}
		if err != nil {
			return nil, nil, err
		}
		helper.Infof("[gRPC] connected to server: %s, addr: %s, timeout: %ds", service.Name, service.Addr, service.Timeout.Seconds)
	}
	cleanup := func() {
		helper.Info("closing the data resources")
	}
	return instance, cleanup, nil
}

func newWarehouseClient(addr string, timeout time.Duration) (warehouse.WarehouseServiceClient, error) {
	conn, err := data.NewClientConnection(addr, timeout)
	return warehouse.NewWarehouseServiceClient(conn), err
}

func newRobotClient(addr string, timeout time.Duration) (robot.RobotServiceClient, error) {
	conn, err := data.NewClientConnection(addr, timeout)
	return robot.NewRobotServiceClient(conn), err
}

func newPlanClient(addr string, timeout time.Duration) (plan.PlanServiceClient, error) {
	conn, err := data.NewClientConnection(addr, timeout)
	return plan.NewPlanServiceClient(conn), err
}
