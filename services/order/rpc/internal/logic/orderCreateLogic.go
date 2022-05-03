package logic

import (
	"context"

	"cinema-shop/services/order/rpc/internal/svc"
	"cinema-shop/services/order/rpc/pb/order"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type OrderCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建订单
func (l *OrderCreateLogic) OrderCreate(in *order.OrderRequest) (*order.OrderResponse, error) {

	resp := new(order.OrderResponse)
	redisObj := redis.New(l.svcCtx.Config.CacheRedis[0].RedisConf.Host)

	script := `
local num = redis.call("SCARD",KEYS[1])
local total_num = ARGV[1]
local buy_num = ARGV[2]
if (num + buy_num) > tonumber(total_num) then
	return "0"
end

for i = 1, buy_num do
	local if_exit = redis.call("SISMEMBER",KEYS[1],ARGV[2+i])
	if if_exit > 0 then
		return "0"
	end
end

for i = 1, buy_num do
	redis.call("SADD",KEYS[1],ARGV[2+i])
end
`

	args := make([]interface{}, 2+len(in.SeatMap))
	args[0] = in.SeatNum
	args[1] = len(in.SeatMap)
	for k, v := range in.SeatMap {
		args[2+k] = v
	}
	result, err := redisObj.Eval(script, []string{"seat:1"}, args...)

	if err != nil && !errors.Is(err, redis.Nil) {
		return resp, errors.Wrap(err, "Order RPC:OrderCreate [Redis Eval] DbError")
	}

	if result != nil {
		return resp, errors.Wrap(err, "Buy Ticket Fail")
	}

	return &order.OrderResponse{
		Ack: 1,
	}, nil
}
