package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"cinema-shop/common/global"
	"cinema-shop/common/utilsx"
	"cinema-shop/services/order/model/orders"
	"cinema-shop/services/order/model/tickets"
	"cinema-shop/services/order/rpc/internal/svc"
	"cinema-shop/services/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type OrderGenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type OrderMessage struct {
	ScreenId int64    `json:"screen_id"`
	Uid      int64    `json:"uid"`
	Seat     []string `json:"seat"`
	Amount   int64    `json:"amount"`
	OrderKey string   `json:"order_key"`
}

type Orders struct {
	OrderId   int64  `db:"order_id"`   // 订单ID
	OrderSn   string `db:"order_sn"`   // 订单编号
	ScreenId  int64  `db:"screen_id"`  // 排片ID
	CreatedAt string `db:"created_at"` // 创建时间
	Uid       int64  `db:"uid"`        // 用户ID
	UpdatedAt string `db:"updated_at"` // 更新时间
	Amount    int64  `db:"amount"`     // 订单金额
	PayTime   string `db:"pay_time"`   // 支付时间
	Status    int64  `db:"status"`     // 订单状态【1待支付 2支付完成 3检票成功 4已退票 5自动取消 6已过期】
	OrderKey  string `db:"order_key"`  // 队列标识
}

func NewOrderGenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderGenLogic {
	return &OrderGenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成订单数据
func (l *OrderGenLogic) OrderGen(in *order.OrderGenRequest) (*order.OrderGenResponse, error) {

	fmt.Println("&&&&&&&&&&&&&&&&&&")
	var orderMsg OrderMessage

	err := json.Unmarshal([]byte(in.Data), &orderMsg)
	if err != nil {
		return &order.OrderGenResponse{}, err
	}
	fmt.Println("************")
	fmt.Println(orderMsg)
	err = l.svcCtx.OrdersModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		order := new(orders.Orders)
		order.Amount = sql.NullInt64{
			Int64: orderMsg.Amount,
			Valid: true,
		}
		order.OrderSn = sql.NullString{
			String: l.genOrderSn(),
			Valid:  true,
		}
		order.ScreenId = sql.NullInt64{
			Int64: orderMsg.ScreenId,
			Valid: true,
		}
		order.Uid = sql.NullInt64{
			Int64: orderMsg.Uid,
			Valid: true,
		}
		order.Status = sql.NullInt64{
			Int64: global.ORDER_CREATE,
			Valid: true,
		}
		order.OrderKey = sql.NullString{
			String: orderMsg.OrderKey,
			Valid:  true,
		}

		order.CreatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}

		orderResult, err := l.svcCtx.OrdersModel.Insert(context, session, order)
		if err != nil {
			return err
		}

		order_id, err := orderResult.LastInsertId()
		if err != nil {
			return err
		}

		ticket := new(tickets.Tickets)
		ticket.OrderId = sql.NullInt64{
			Int64: order_id,
			Valid: true,
		}
		ticket.ScreenId = sql.NullInt64{
			Int64: 1,
			Valid: true,
		}
		ticket.TicketSn = sql.NullString{
			String: l.genTicketSn(),
			Valid:  true,
		}

		ticket.Status = sql.NullInt64{
			Int64: global.TICKET_CREATE,
			Valid: true,
		}
		ticket.CreatedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}

		for _, v := range orderMsg.Seat {
			ticket.Seat = sql.NullString{
				String: v,
				Valid:  true,
			}
			l.svcCtx.TicketsModel.Insert(context, session, ticket)
		}

		return nil
	})
	if err != nil {
		return &order.OrderGenResponse{}, nil
	}

	return &order.OrderGenResponse{}, nil
}

func (l *OrderGenLogic) genOrderSn() string {

	return fmt.Sprintf("OR%s%s", time.Now().Format("20060102"), utilsx.RandString(6))
}

func (l *OrderGenLogic) genTicketSn() string {

	return fmt.Sprintf("TICKET%s", utilsx.RandString(12))
}
