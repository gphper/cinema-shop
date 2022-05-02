package Screen

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"cinema-shop/services/cinema/rpc/cinema"
	"cinema-shop/services/order/rpc/order"

	cuserror "github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"golang.org/x/sync/singleflight"
)

type ScreenDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var gsf singleflight.Group

var SEATEMPTYERR = errors.New("empty seat")

func NewScreenDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenDetailLogic {
	return &ScreenDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScreenDetailLogic) ScreenDetail(req *types.ScreenDetailReq) (resp *types.ScreenDetailResp, err error) {

	resp = new(types.ScreenDetailResp)
	screenDetailResp, err := l.svcCtx.CinemaRpcClient.ScreenDetail(l.ctx, &cinema.ScreenDetailRequest{
		ScreenId: req.ScreenId,
	})

	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Get Screen Data Fail")
	}

	//从缓存中获取已被占座位
	soldSeat, err := l._GetSoldSeat(req.ScreenId)
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Get Sold Seat Fail")
	}

	var seatMap [][]int
	err = json.Unmarshal([]byte(screenDetailResp.CurrentSeat), &seatMap)
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Unmarshal Error")
	}

	for _, v := range soldSeat {
		seatArr := strings.Split(v, "#")
		if seatArr[0] == "init" {
			continue
		}
		x, err := strconv.Atoi(seatArr[0])
		if err != nil {
			return resp, errorxx.NewCodeError(2001, "Base Error")
		}

		y, err := strconv.Atoi(seatArr[1])
		if err != nil {
			return resp, errorxx.NewCodeError(2001, "Base Error")
		}

		seatMap[x][y] = 2
	}

	seatMapByte, err := json.Marshal(seatMap)
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Marshal Error")
	}

	resp.Data.ScreenId = screenDetailResp.ScreenId
	resp.Data.SeatMap = string(seatMapByte)

	return
}

func (l *ScreenDetailLogic) _GetSoldSeat(screenId int64) ([]string, error) {

	key := fmt.Sprintf("seat:%d", screenId)
	seats, err := l._GetCache(key)

	if errors.Is(err, SEATEMPTYERR) {
		v, err, _ := gsf.Do(key, func() (interface{}, error) {
			dbSeat, err := l._GetDb(screenId)
			if err != nil {
				return nil, err
			}
			if len(dbSeat) == 0 {
				err := l._SetCache(key, []string{"init#init"})
				if err != nil {
					return nil, err
				}
				return []string{"init#init"}, nil
			}

			err = l._SetCache(key, dbSeat)
			if err != nil {
				return nil, err
			}
			return dbSeat, nil
		})
		if err != nil {

			return nil, err
		}

		return v.([]string), nil
	} else if err != nil {

		return nil, err
	}

	return seats, nil
}

func (l *ScreenDetailLogic) _GetCache(key string) ([]string, error) {
	//redis操作座位
	redisObj := redis.New(l.svcCtx.Config.CacheRedis[0].RedisConf.Host)
	seatMap, err := redisObj.SmembersCtx(l.ctx, key)
	if err != nil {
		return nil, errorxx.NewCodeError(2001, "Get Stock Fail")
	}
	if len(seatMap) == 0 {
		return nil, SEATEMPTYERR
	}
	return seatMap, nil
}

func (l *ScreenDetailLogic) _GetDb(screenId int64) ([]string, error) {
	//获取已售出的座位信息
	ticketResp, err := l.svcCtx.OrderRpcClient.TicketSeat(l.ctx, &order.TicketSeatRequest{
		ScreenId: screenId,
	})

	if err != nil {
		return nil, cuserror.Wrap(err, "")
	}

	return ticketResp.Seat, nil
}

func (l *ScreenDetailLogic) _SetCache(key string, values []string) error {

	redisObj := redis.New(l.svcCtx.Config.CacheRedis[0].RedisConf.Host)
	_, err := redisObj.SaddCtx(l.ctx, key, values)
	if err != nil {

		return err
	}

	return nil
}
