package Film

import (
	"context"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilmListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmListLogic {
	return &FilmListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilmListLogic) FilmList(req *types.FilmListReq) (resp *types.FilmListResp, err error) {
	filmListRpcResp, err := l.svcCtx.CinemaRpcClient.List(l.ctx, &cinema.FilmListRequest{
		Page:  int32(req.Page),
		Limit: int32(req.Limit),
	})
	if err != nil {
		return resp, errorxx.NewCodeError(2001, err.Error())
	}

	var infos []types.FilmListInfo
	if len(filmListRpcResp.Data) > 0 {
		infos = make([]types.FilmListInfo, len(filmListRpcResp.Data))
		for k, v := range filmListRpcResp.Data {
			infos[k] = types.FilmListInfo{
				FilmId:   int(v.GetFilmId()),
				FilmName: v.GetFilmName(),
				Type:     int(v.GetType()),
				CoverPic: v.GetCoverPic(),
			}
		}
	}

	resp = &types.FilmListResp{
		Data: types.FilmListInfoResp{
			Data:  infos,
			Count: int(filmListRpcResp.Count),
			Page:  int(filmListRpcResp.Page),
		},
	}

	return
}
