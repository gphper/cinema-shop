package Film

import (
	"context"

	"cinema-shop/services/cinema/api/internal/svc"
	"cinema-shop/services/cinema/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
