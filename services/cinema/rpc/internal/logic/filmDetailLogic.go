package logic

import (
	"context"

	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilmDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFilmDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmDetailLogic {
	return &FilmDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 影片详情
func (l *FilmDetailLogic) FilmDetail(in *cinema.FilmDatailRequest) (*cinema.FilmDetailResponse, error) {
	filmInfo, err := l.svcCtx.FilmModel.FindOne(l.ctx, int64(in.FilmId))
	if err != nil {
		return &cinema.FilmDetailResponse{}, err
	}

	return &cinema.FilmDetailResponse{
		Data: &cinema.FilmDetailInfo{
			FilmId:   int32(filmInfo.FilmId),
			Cate:     int32(filmInfo.Cate),
			FilmName: filmInfo.FilmName,
			CoverPic: filmInfo.CoverPic,
			FilmDesc: filmInfo.FilmDesc,
		},
	}, nil
}
