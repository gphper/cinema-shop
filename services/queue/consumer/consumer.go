package main

import (
	"cinema-shop/services/order/rpc/order"
	"cinema-shop/services/queue/consumer/config"
	"cinema-shop/services/queue/rabbitmq"
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "./etc/consumer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	mqUrl := fmt.Sprintf("amqp://%s:%s@%s/", c.RabbitMq.Username, c.RabbitMq.Password, c.RabbitMq.Host)
	rabbit := *rabbitmq.NewRabbitMQ(mqUrl)
	msgChan, err := rabbit.Channel.Consume(
		"order_create_queue", // queue
		"",                   // consumer
		false,                // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		panic(err.Error())
	}

	forever := make(chan bool)

	orderRpc := order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf))
	go func() {
		for d := range msgChan {
			log.Printf("收到消息: %s", d.Body)

			//处理生成订单操作
			_, err := orderRpc.OrderGen(context.Background(), &order.OrderGenRequest{
				Data: string(d.Body),
			})
			if err != nil {
				fmt.Println(err)
			}

			// d.Ack(true)
		}
	}()

	log.Printf("wait message")
	<-forever
}
