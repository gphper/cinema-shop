package logic

import (
	"context"

	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FilmAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFilmAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmAllLogic {
	return &FilmAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据影片ID获取全部影片信息
func (l *FilmAllLogic) FilmAll(in *cinema.FilmAllRequest) (*cinema.FilmAllResponse, error) {

	resp := new(cinema.FilmAllResponse)

	sqlBuilder := l.svcCtx.FilmModel.RowBuilder()
	sqlBuilder = sqlBuilder.Where(squirrel.Eq{"film_id": in.FilmId})
	filmInfos, err := l.svcCtx.FilmModel.FindAll(l.ctx, sqlBuilder, "")
	if err != nil {
		return &cinema.FilmAllResponse{}, errors.Wrap(err, "Cinema RPC:FilmAll DbError")
	}

	copier.Copy(&resp.Data, &filmInfos)

	return resp, nil
}
