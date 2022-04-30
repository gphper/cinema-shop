package logic

import (
	"context"

	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScreenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScreenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenListLogic {
	return &ScreenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据日期、影院ID、影片ID获取排片场次
func (l *ScreenListLogic) ScreenList(in *cinema.ScreenListRequest) (*cinema.ScreenListResponse, error) {

	resp := new(cinema.ScreenListResponse)
	sqlBuilder := l.svcCtx.ScreenModel.RowBuilder()
	sqlBuilder = sqlBuilder.Where(squirrel.Eq{"t_date": in.TDate}).Where(squirrel.Eq{"film_id": in.FilmId}).Where(squirrel.Eq{"cinema_id": in.CinemaId})
	screenInfos, err := l.svcCtx.ScreenModel.FindAll(l.ctx, sqlBuilder, "start_time asc")
	if err != nil {
		return resp, errors.Wrap(err, "Cinema Rpc:ScreenList DbError [ScreenModel.FindAll]")
	}

	copier.Copy(&resp.Data, &screenInfos)

	return resp, nil
}
