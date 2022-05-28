package Order

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/cinema/rpc/cinema"
	"cinema-shop/services/order/api/internal/svc"
	"cinema-shop/services/order/api/internal/types"
	"cinema-shop/services/order/rpc/order"
	"cinema-shop/services/queue/rpc/queue"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req *types.OrderReq) (resp *types.OrderResp, err error) {

	resp = new(types.OrderResp)

	id, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Get UserId Fail")
	}

	screenResp, err := l.svcCtx.CinemaRpcClient.ScreenDetail(l.ctx, &cinema.ScreenDetailRequest{
		ScreenId: req.ScreenId,
	})
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "CinemaRpcClient ScreenDetail Error")
	}

	var seatMap [][]int
	err = json.Unmarshal([]byte(screenResp.CurrentSeat), &seatMap)
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Unmarshal Error")
	}

	row := len(seatMap)
	col := len(seatMap[0])

	if len(req.SeatMap) > 0 {
		for _, v := range req.SeatMap {
			seatArr := strings.Split(v, "#")
			x, _ := strconv.Atoi(seatArr[0])
			y, _ := strconv.Atoi(seatArr[1])
			if x > row-1 || y > col-1 || seatMap[x][y] == 0 {
				return resp, errorxx.NewCodeError(20010, "Illegal Seat Information")
			}
		}
	}

	_, err = l.svcCtx.OrderRpcClient.OrderCreate(l.ctx, &order.OrderRequest{
		ScreenId: req.ScreenId,
		SeatMap:  req.SeatMap,
		SeatNum:  screenResp.SeatNum,
	})

	if err != nil {
		return resp, errorxx.NewCodeError(2001, err.Error())
	}

	queueResp, err := l.svcCtx.QueueRpcClient.OrderQueue(l.ctx, &queue.OrderCreateRequest{
		ScreenId: req.ScreenId,
		Uid:      id,
		Seat:     req.SeatMap,
		Amount:   screenResp.Price * int64(len(req.SeatMap)),
	})
	if err != nil {
		return resp, errorxx.NewCodeError(2001, err.Error())
	}

	resp.OrderKey = queueResp.OrderKey

	return
}
