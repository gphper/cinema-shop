package Screen

import (
	"context"
	"fmt"

	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"cinema-shop/services/cinema/rpc/cinema"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScreenCinemaInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScreenCinemaInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenCinemaInfoLogic {
	return &ScreenCinemaInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScreenCinemaInfoLogic) ScreenCinemaInfo(req *types.ScreenCinemaInfoReq) (resp *types.ScreenCinemaInfoResp, err error) {

	var cinemaMap map[int32]cinema.CinemaInfo
	var cinemaIds []int32
	var data []types.ScreenCinemaInfo

	cinemaInfoResp, err := l.svcCtx.CinemaRpcClient.CinemaInfo(l.ctx, &cinema.CinemaInfoRequest{
		CityCode: req.CityCode,
		AreaCode: req.AreaCode,
	})
	if err != nil {
		err = errors.Wrap(err, "Cinema Api ScreenCinemaInfo Error")
		return
	}

	if len(cinemaInfoResp.Data) > 0 {
		cinemaMap = make(map[int32]cinema.CinemaInfo, len(cinemaInfoResp.Data))
		cinemaIds = make([]int32, len(cinemaInfoResp.Data))

		for k, v := range cinemaInfoResp.Data {
			cinemaMap[v.CinemaId] = cinema.CinemaInfo{
				CinemaName: v.CinemaName,
				Place:      v.Place,
				Score:      v.Score,
				Tags:       v.Tags,
			}
			cinemaIds[k] = v.CinemaId
		}
	}

	screenCinemaInfoResp, err := l.svcCtx.CinemaRpcClient.ScreenCinemaInfo(l.ctx, &cinema.ScreenCinemaInfoRequest{
		FilmId:    int32(req.FilmId),
		CinemaIds: cinemaIds,
		TDate:     req.TDate,
	})
	if err != nil {
		err = errors.Wrap(err, "Cinema Api ScreenCinemaInfo Error")
		return
	}

	fmt.Println(screenCinemaInfoResp.Data)
	if len(screenCinemaInfoResp.Data) > 0 {
		data = make([]types.ScreenCinemaInfo, len(screenCinemaInfoResp.Data))
		for k, v := range screenCinemaInfoResp.Data {
			tempCinema, ok := cinemaMap[v.CinemaId]
			if ok {
				data[k] = types.ScreenCinemaInfo{
					Price:      int(v.Price),
					Film:       v.StartTime,
					CinemaName: tempCinema.CinemaName,
					Place:      tempCinema.Place,
					Score:      int(tempCinema.Score),
					Tags:       tempCinema.Tags,
				}
			}
		}
	}

	resp = &types.ScreenCinemaInfoResp{
		Data: data,
	}

	return
}
