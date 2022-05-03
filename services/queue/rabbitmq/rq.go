package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
) //导入mq包

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	//MQ链接字符串
	Mqurl string
}

// 创建结构体实例
func NewRabbitMQ(mqUrl string) *RabbitMQ {
	rabbitMQ := RabbitMQ{
		Mqurl: mqUrl,
	}
	var err error
	//创建rabbitmq连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	checkErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	checkErr(err, "创建channel失败")

	return &rabbitMQ

}

//初始化exchange，queue，key
func (mq *RabbitMQ) Init(queueName, exchange, routingKey string) {
	//初始化队列excahnge
	_, err := mq.Channel.QueueDeclare( // 返回的队列对象内部记录了队列的一些信息，这里没什么用
		queueName, // 队列名
		true,      // 是否持久化
		false,     // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false,     // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false,     // 是否阻塞
		nil,       //额外属性（我还不会用）
	)
	if err != nil {
		panic("声明队列失败")
	}

	// 2.声明交换器
	err = mq.Channel.ExchangeDeclare(
		exchange, //交换器名
		"topic",  //exchange type：一般用fanout、direct、topic
		true,     // 是否持久化
		false,    //是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
		false,    //设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
		false,    // 是否阻塞
		nil,      // 额外属性
	)
	if err != nil {
		panic("声明交换器失败")
	}

	// 3.建立Binding(可随心所欲建立多个绑定关系)
	err = mq.Channel.QueueBind(
		queueName,  // 绑定的队列名称
		routingKey, // bindkey 用于消息路由分发的key
		exchange,   // 绑定的exchange名
		false,      // 是否阻塞
		nil,        // 额外属性
	)

	if err != nil {
		panic("绑定队列和交换器失败")
	}
}

// 释放资源,建议NewRabbitMQ获取实例后 配合defer使用
func (mq *RabbitMQ) ReleaseRes() {
	mq.Conn.Close()
	mq.Channel.Close()
}

func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}
