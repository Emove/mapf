package event

// Configuration 事件中心配置
type Configuration interface {
	GetAddr() string
}
