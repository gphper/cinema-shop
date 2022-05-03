package svc

import (
	"cinema-shop/services/cinema/rpc/cinema"
	"cinema-shop/services/order/api/internal/config"
	"cinema-shop/services/order/rpc/order"
	"cinema-shop/services/queue/rpc/queue"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	CinemaRpcClient cinema.Cinema
	OrderRpcClient  order.Order
	QueueRpcClient  queue.Queue
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CinemaRpcClient: cinema.NewCinema(zrpc.MustNewClient(c.CinemaRpcConf)),
		OrderRpcClient:  order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		QueueRpcClient:  queue.NewQueue(zrpc.MustNewClient(c.QueueRpcConf)),
	}
}
