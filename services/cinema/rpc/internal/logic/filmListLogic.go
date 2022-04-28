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

type FilmListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFilmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmListLogic {
	return &FilmListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 影片列表
func (l *FilmListLogic) FilmList(in *cinema.FilmListRequest) (*cinema.FilmListResponse, error) {

	sqlBuilder := l.svcCtx.FilmModel.RowBuilder()
	countBuilder := l.svcCtx.FilmModel.CountBuilder("film_id")
	if in.Cate > 0 {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"cate": in.Cate})
		countBuilder = countBuilder.Where(squirrel.Eq{"cate": in.Cate})
	}

	if in.Type > 0 {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"type": in.Type})
		countBuilder = countBuilder.Where(squirrel.Eq{"type": in.Type})
	}

	if in.Status > 0 {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"status": in.Status})
		countBuilder = countBuilder.Where(squirrel.Eq{"status": in.Status})
	}

	data, err := l.svcCtx.FilmModel.FindPageListByPage(l.ctx, sqlBuilder, int64(in.Page), int64(in.Limit), "")
	if err != nil {
		return &cinema.FilmListResponse{}, errors.Wrap(err, "Cinema RPC:FilmList Error")
	}

	var datas []*cinema.FilmListInfo
	copier.Copy(&datas, data)

	count, err := l.svcCtx.FilmModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return &cinema.FilmListResponse{}, errors.Wrap(err, "Cinema RPC:FilmList Error")
	}

	return &cinema.FilmListResponse{
		Data:  datas,
		Count: int32(count),
		Page:  in.Page,
	}, nil
}
