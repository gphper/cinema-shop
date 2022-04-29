package logic

import (
	"context"
	"time"

	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScreenCinemaInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScreenCinemaInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenCinemaInfoLogic {
	return &ScreenCinemaInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据影片和日期和影院ID获取排片信息
func (l *ScreenCinemaInfoLogic) ScreenCinemaInfo(in *cinema.ScreenCinemaInfoRequest) (*cinema.ScreenCinemaInfoResponse, error) {

	sqlBuilder := l.svcCtx.ScreenModel.RowCusBuilder("cinema_id,group_concat(start_time ORDER BY start_time ASC) as film,MIN(price) price")
	sqlBuilder = sqlBuilder.Where(squirrel.Eq{"cinema_id": in.CinemaIds})
	sqlBuilder = sqlBuilder.Where(squirrel.GtOrEq{"start_time": time.Now().Format("15:04:05")})

	data, err := l.svcCtx.ScreenModel.FindAllS(l.ctx, sqlBuilder, "screen_id ASC")
	if err != nil {
		return &cinema.ScreenCinemaInfoResponse{}, errors.Wrap(err, "Cinema Rpc:ScreenCinemaInfo Error")
	}

	var info []*cinema.ScreenCinemaInfo
	copier.Copy(&info, &data)

	return &cinema.ScreenCinemaInfoResponse{
		Data: info,
	}, nil
}
