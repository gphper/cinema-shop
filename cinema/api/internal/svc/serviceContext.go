package svc

import (
	"cinema-shop/services/film/api/internal/config"
	"cinema-shop/services/film/rpc/film"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	FilmRpcClient film.Film
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		FilmRpcClient: film.NewFilm(zrpc.MustNewClient(c.FilmRpcConf)),
	}
}
