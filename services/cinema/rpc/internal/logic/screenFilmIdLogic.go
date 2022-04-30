package logic

import (
	"context"

	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScreenFilmIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScreenFilmIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenFilmIdLogic {
	return &ScreenFilmIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据日期、影院ID获取排片电影
func (l *ScreenFilmIdLogic) ScreenFilmId(in *cinema.ScreenFilmIdRequest) (*cinema.ScreenFilmIdResponse, error) {

	resp := new(cinema.ScreenFilmIdResponse)
	sqlBuilder := l.svcCtx.ScreenModel.RowBuilder()
	sqlBuilder = sqlBuilder.Where(squirrel.Eq{"cinema_id": in.CinemaId}).Where(squirrel.GtOrEq{"t_date": in.TDate})
	result, err := l.svcCtx.ScreenModel.FindAll(l.ctx, sqlBuilder, "")
	if err != nil {
		return resp, errors.Wrap(err, "Cinema Rpc:ScreenFilmId DbError")
	}

	if len(result) > 0 {
		filter := make(map[int64]struct{}, len(result))
		resp.FilmId = make([]int64, 0, len(result))
		for _, screen := range result {
			_, ok := filter[screen.FilmId.Int64]
			if !ok {
				filter[screen.FilmId.Int64] = struct{}{}
				resp.FilmId = append(resp.FilmId, screen.FilmId.Int64)
			}
		}
	}

	return resp, nil
}
