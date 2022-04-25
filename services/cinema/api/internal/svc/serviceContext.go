package svc

import (
	"cinema-shop/services/cinema/api/internal/config"
	"cinema-shop/services/cinema/rpc/cinema"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	CinemaRpcClient cinema.Cinema
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CinemaRpcClient: cinema.NewCinema(zrpc.MustNewClient(c.CinemaRpcConf)),
	}
}
