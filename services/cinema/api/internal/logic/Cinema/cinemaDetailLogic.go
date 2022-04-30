package Cinema

import (
	"context"
	"time"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"cinema-shop/services/cinema/rpc/cinema"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CinemaDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCinemaDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CinemaDetailLogic {
	return &CinemaDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CinemaDetailLogic) CinemaDetail(req *types.CinemaDetailReq) (resp *types.CinemaDatailResp, err error) {

	resp = new(types.CinemaDatailResp)

	cinemaDetailResp, err := l.svcCtx.CinemaRpcClient.CinemaDetail(l.ctx, &cinema.CinemaDetailRequest{
		CinemaId: req.CinemaId,
	})
	if err != nil {
		return &types.CinemaDatailResp{}, errorxx.NewCodeError(2001, err.Error())
	}

	filmIdResp, err := l.svcCtx.CinemaRpcClient.ScreenFilmId(l.ctx, &cinema.ScreenFilmIdRequest{
		TDate:    time.Now().Format("2006-01-02"),
		CinemaId: req.CinemaId,
	})
	if err != nil {
		return &types.CinemaDatailResp{}, errorxx.NewCodeError(2001, err.Error())
	}

	filmInfoResp, err := l.svcCtx.CinemaRpcClient.FilmAll(l.ctx, &cinema.FilmAllRequest{
		FilmId: filmIdResp.FilmId,
	})
	if err != nil {
		return &types.CinemaDatailResp{}, errorxx.NewCodeError(2001, err.Error())
	}

	copier.Copy(&resp.Data, &cinemaDetailResp)
	filmInfos := make([]types.FilmInfo, 0)
	copier.Copy(&filmInfos, &filmInfoResp.Data)
	resp.Data.FilmInfos = filmInfos

	return
}
