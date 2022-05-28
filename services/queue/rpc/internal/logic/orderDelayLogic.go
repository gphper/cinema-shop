package logic

import (
	"context"
	"encoding/json"

	"cinema-shop/services/queue/rpc/internal/svc"
	"cinema-shop/services/queue/rpc/pb/queue"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDelayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDelayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDelayLogic {
	return &OrderDelayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自动取消未支付订单队列
func (l *OrderDelayLogic) OrderDelay(in *queue.OrderDelayRequest) (*queue.OrderDelayResponse, error) {
	var (
		exchange   = "exchange_order_delay"
		routingKey = "order_delay"
	)

	type OrderMessage struct {
		OrderId int64 `json:"order_id"`
	}

	orderMessage := OrderMessage{
		OrderId: in.OrderId,
	}

	msgJson, err := json.Marshal(orderMessage)
	if err != nil {
		return &queue.OrderDelayResponse{}, errors.Wrap(err, "Queue RPC:OrderQueue [Marshal Err] Error")
	}

	headers := make(amqp.Table)
	headers["x-delay"] = 600000

	err = l.svcCtx.OrderRabbitMq.Channel.Publish(
		exchange,   // 交换器名
		routingKey, // routing key
		false,      // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
		false,      // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
		amqp.Publishing{ // 发送的消息，固定有消息体和一些额外的消息头，包中提供了封装对象
			ContentType: "application/json", // 消息内容的类型
			Body:        msgJson,            // 消息内容
			Headers:     headers,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "Queue RPC:OrderQueue [Publish Message Fail] Error")
	}

	return &queue.OrderDelayResponse{
		Ack: 1,
	}, nil
}
