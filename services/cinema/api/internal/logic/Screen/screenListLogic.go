package Screen

import (
	"context"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"cinema-shop/services/cinema/rpc/cinema"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScreenListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScreenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenListLogic {
	return &ScreenListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScreenListLogic) ScreenList(req *types.ScreenListReq) (resp *types.ScreenListResp, err error) {

	resp = new(types.ScreenListResp)
	var hallMap map[int64]string
	hallListResp, err := l.svcCtx.CinemaRpcClient.HallList(l.ctx, &cinema.HallListRequest{
		CinemaId: req.CinemaId,
	})
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Get Cinema:HallList Error")
	}

	if len(hallListResp.Data) > 0 {
		hallMap = make(map[int64]string, len(hallListResp.Data))
		for _, hall := range hallListResp.Data {
			hallMap[hall.HallId] = hall.HallName
		}
	}

	screenListResp, err := l.svcCtx.CinemaRpcClient.ScreenList(l.ctx, &cinema.ScreenListRequest{
		TDate:    req.TDate,
		CinemaId: req.CinemaId,
		FilmId:   req.FilmId,
	})
	if err != nil {
		return resp, errorxx.NewCodeError(2001, "Get Cinema:ScreenList Error")
	}

	if len(screenListResp.Data) > 0 {
		resp.Data = make([]types.ScreenInfo, len(screenListResp.Data))
		for k, v := range screenListResp.Data {
			resp.Data[k] = types.ScreenInfo{
				ScreenId:  v.ScreenId,
				StartTime: v.StartTime,
				Price:     v.Price,
				HallName:  hallMap[v.HallId],
			}
		}
	}

	return
}
