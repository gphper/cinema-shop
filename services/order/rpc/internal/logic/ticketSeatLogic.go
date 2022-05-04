package logic

import (
	"context"

	"cinema-shop/services/order/rpc/internal/svc"
	"cinema-shop/services/order/rpc/pb/order"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type TicketSeatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}



func NewTicketSeatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketSeatLogic {
	return &TicketSeatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据排片ID获取已被占用的座位
func (l *TicketSeatLogic) TicketSeat(in *order.TicketSeatRequest) (*order.TicketSeatResponse, error) {

	resp := new(order.TicketSeatResponse)
	var seats []string

	sqlBuilder := l.svcCtx.TicketsModel.RowBuilder().Where(squirrel.Eq{"screen_id": in.ScreenId})
	ticketSeat, err := l.svcCtx.TicketsModel.FindAll(l.ctx, sqlBuilder, "ticket_id asc")
	if err != nil {
		return resp, errors.Wrap(err, "Order RPC:TicketSeat [TicketsModel.FindAll] DbError")
	}

	if len(ticketSeat) > 0 {
		seats = make([]string, len(ticketSeat))
		for k, v := range ticketSeat {
			seats[k] = v.Seat.String
		}
	}

	resp.Seat = seats

	return resp, nil
}
