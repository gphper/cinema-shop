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

type CinemaInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCinemaInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CinemaInfoLogic {
	return &CinemaInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据地理位置获取影院信息
func (l *CinemaInfoLogic) CinemaInfo(in *cinema.CinemaInfoRequest) (*cinema.CinemaInfoResponse, error) {

	sqlBuilder := l.svcCtx.CinemaModel.RowBuilder()

	if in.CityCode != "" {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"city": in.CityCode})
	}

	if in.AreaCode != "" {
		sqlBuilder = sqlBuilder.Where(squirrel.Eq{"area": in.AreaCode})
	}

	cinemaData, err := l.svcCtx.CinemaModel.FindAll(l.ctx, sqlBuilder, "cinema_id ASC")
	if err != nil {
		return &cinema.CinemaInfoResponse{}, errors.Wrap(err, "Cinema RPC:CinemaInfo Error")
	}

	var resp []*cinema.CinemaInfo
	copier.Copy(&resp, &cinemaData)

	return &cinema.CinemaInfoResponse{
		Data: resp,
	}, nil
}
