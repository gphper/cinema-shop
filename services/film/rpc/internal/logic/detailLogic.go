package logic

import (
	"context"

	"cinema-shop/services/film/rpc/internal/svc"
	"cinema-shop/services/film/rpc/pb/film"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 影片详情
func (l *DetailLogic) Detail(in *film.FilmDatailRequest) (*film.FilmDetailResponse, error) {
	// todo: add your logic here and delete this line

	return &film.FilmDetailResponse{}, nil
}
