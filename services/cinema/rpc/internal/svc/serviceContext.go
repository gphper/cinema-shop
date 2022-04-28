package svc

import (
	"cinema-shop/services/cinema/model/cinema"
	"cinema-shop/services/cinema/model/film"
	"cinema-shop/services/cinema/model/screen"
	"cinema-shop/services/cinema/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	FilmModel   film.FilmModel
	CinemaModel cinema.CinemaModel
	ScreenModel screen.ScreenModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		FilmModel:   film.NewFilmModel(conn, c.CacheRedis),
		CinemaModel: cinema.NewCinemaModel(conn, c.CacheRedis),
		ScreenModel: screen.NewScreenModel(conn, c.CacheRedis),
	}
}
