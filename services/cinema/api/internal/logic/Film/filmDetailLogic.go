package Film

import (
	"context"

	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"cinema-shop/services/cinema/rpc/cinema"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilmDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmDetailLogic {
	return &FilmDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilmDetailLogic) FilmDetail(req *types.FilmDetailReq) (resp *types.FilmDetailResp, err error) {

	rpcResp, err := l.svcCtx.CinemaRpcClient.FilmDetail(l.ctx, &cinema.FilmDatailRequest{
		FilmId: int32(req.FilmId),
	})
	if err != nil {
		return resp, err
	}

	resp = &types.FilmDetailResp{
		Data: types.FilmDetailInfo{
			FilmId:   int(rpcResp.Data.FilmId),
			FilmName: rpcResp.Data.FilmName,
			Cate:     int(rpcResp.Data.Cate),
			CoverPic: rpcResp.Data.CoverPic,
			FilmDesc: rpcResp.Data.FilmDesc,
		},
	}

	return
}
