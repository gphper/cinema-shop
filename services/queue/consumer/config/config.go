package config

import "github.com/zeromicro/go-zero/zrpc"

type RabbitMqConf struct {
	Host     string `json:",optional"`
	Port     string `json:",optional"`
	Username string `json:",optional"`
	Password string `json:",optional"`
}

type Config struct {
	RabbitMq     RabbitMqConf
	OrderRpcConf zrpc.RpcClientConf
	QueueRpcConf zrpc.RpcClientConf
}
