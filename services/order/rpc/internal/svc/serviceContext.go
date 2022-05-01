package svc

import (
	"cinema-shop/services/order/model/tickets"
	"cinema-shop/services/order/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	TicketsModel tickets.TicketsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		TicketsModel: tickets.NewTicketsModel(conn, c.CacheRedis),
	}
}
