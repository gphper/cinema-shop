package svc

import (
	"cinema-shop/services/usercenter/model/user"
	"cinema-shop/services/usercenter/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(conn, c.CacheRedis),
	}
}
