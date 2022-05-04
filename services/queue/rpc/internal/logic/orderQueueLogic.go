package logic

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"cinema-shop/services/queue/rpc/internal/svc"
	"cinema-shop/services/queue/rpc/pb/queue"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderQueueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderQueueLogic {

	return &OrderQueueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建订单队列
func (l *OrderQueueLogic) OrderQueue(in *queue.OrderCreateRequest) (*queue.OrderCreateResponse, error) {

	var (
		exchange   = "exchange_order"
		routingKey = "order_create"
	)

	type OrderMessage struct {
		ScreenId int64    `json:"screen_id"`
		Uid      int64    `json:"uid"`
		Amount   int64    `json:"amount"`
		Seat     []string `json:"seat"`
		OrderKey string   `json:"order_key"`
	}

	orderMessage := OrderMessage{
		ScreenId: in.ScreenId,
		Uid:      in.Uid,
		Amount:   in.Amount,
		Seat:     in.Seat,
	}

	key := fmt.Sprintf("%d:%d:%s:%s", in.ScreenId, in.Uid, in.Seat, time.Nanosecond)
	orderMessage.OrderKey = fmt.Sprintf("%x", md5.Sum([]byte(key)))

	msgJson, err := json.Marshal(orderMessage)
	if err != nil {
		return &queue.OrderCreateResponse{}, errors.Wrap(err, "Queue RPC:OrderQueue [Marshal Err] Error")
	}

	err = l.svcCtx.OrderRabbitMq.Channel.Publish(
		exchange,   // 交换器名
		routingKey, // routing key
		false,      // 是否返回消息(匹配队列)，如果为true, 会根据binding规则匹配queue，如未匹配queue，则把发送的消息返回给发送者
		false,      // 是否返回消息(匹配消费者)，如果为true, 消息发送到queue后发现没有绑定消费者，则把发送的消息返回给发送者
		amqp.Publishing{ // 发送的消息，固定有消息体和一些额外的消息头，包中提供了封装对象
			ContentType: "application/json", // 消息内容的类型
			Body:        msgJson,            // 消息内容
		},
	)

	if err != nil {
		return nil, errors.Wrap(err, "Queue RPC:OrderQueue [Publish Message Fail] Error")
	}

	return &queue.OrderCreateResponse{
		OrderKey: orderMessage.OrderKey,
	}, nil
}
