package svc

import (
	"cinema-shop/services/film/model/film"
	"cinema-shop/services/film/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	FilmModel film.FilmModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		FilmModel: film.NewFilmModel(conn, c.CacheRedis),
	}
}
