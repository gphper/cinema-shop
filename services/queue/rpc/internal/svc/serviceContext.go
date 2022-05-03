package svc

import (
	"cinema-shop/services/queue/rabbitmq"
	"cinema-shop/services/queue/rpc/internal/config"
	"fmt"
)

type ServiceContext struct {
	Config        config.Config
	OrderRabbitMq rabbitmq.RabbitMQ
}

func NewServiceContext(c config.Config) *ServiceContext {
	mqUrl := fmt.Sprintf("amqp://%s:%s@%s/", c.RabbitMq.Username, c.RabbitMq.Password, c.RabbitMq.Host)
	rabbit := *rabbitmq.NewRabbitMQ(mqUrl)

	rabbit.Init("order_create_queue", "exchange_order", "order_create")
	return &ServiceContext{
		Config:        c,
		OrderRabbitMq: rabbit,
	}
}
