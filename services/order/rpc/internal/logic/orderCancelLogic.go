package logic

import (
	"context"
	"database/sql"
	"fmt"

	"cinema-shop/common/global"
	"cinema-shop/services/order/model/orders"
	"cinema-shop/services/order/model/tickets"
	"cinema-shop/services/order/rpc/internal/svc"
	"cinema-shop/services/order/rpc/pb/order"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type OrderCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCancelLogic {
	return &OrderCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消订单
func (l *OrderCancelLogic) OrderCancel(in *order.OrderCancelRequest) (*order.OrderCancelResponse, error) {

	err := l.svcCtx.OrdersModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		if err := l.svcCtx.OrdersModel.UpdateSome(context, session, &orders.Orders{
			OrderId: in.OrderId,
			Status: sql.NullInt64{
				Int64: global.ORDER_AUTOCANCEL,
				Valid: true,
			},
		}); err != nil {
			return err
		}

		if err := l.svcCtx.TicketsModel.UpdateByOrderId(l.ctx, session, in.OrderId, &tickets.Tickets{
			Status: sql.NullInt64{
				Int64: global.TICKET_AUTOCANCEL,
				Valid: true,
			},
			OrderId: sql.NullInt64{
				Int64: in.OrderId,
				Valid: true,
			},
		}); err != nil {
			return err
		}

		//释放占住的座位
		sqlBuilder := l.svcCtx.TicketsModel.RowBuilder()
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"order_id": in.OrderId})
		tickets, err := l.svcCtx.TicketsModel.FindAll(context, sqlBuilder, "")
		if err != nil {
			return err
		}

		var values []interface{}
		var screenId int64
		if len(tickets) > 0 {
			values = make([]interface{}, len(tickets))
			for k, v := range tickets {
				values[k] = v.Seat.String
				screenId = v.ScreenId.Int64
			}
		}

		redisObj := redis.New(l.svcCtx.Config.CacheRedis[0].RedisConf.Host)
		key := fmt.Sprintf("seat:%d", screenId)

		_, err = redisObj.Srem(key, values...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &order.OrderCancelResponse{
		Ack: "1",
	}, nil
}
