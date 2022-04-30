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

type HallListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHallListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HallListLogic {
	return &HallListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据影院ID获取影厅列表
func (l *HallListLogic) HallList(in *cinema.HallListRequest) (*cinema.HallListResponse, error) {

	resp := new(cinema.HallListResponse)

	sqlBuilder := l.svcCtx.HallModel.RowBuilder()
	sqlBuilder = sqlBuilder.Where(squirrel.Eq{"cinema_id": in.CinemaId})
	hallInfo, err := l.svcCtx.HallModel.FindAll(l.ctx, sqlBuilder, "")
	if err != nil {
		return &cinema.HallListResponse{}, errors.Wrap(err, "Cinema RPC:HallList DbError")
	}

	copier.Copy(&resp.Data, &hallInfo)
	return resp, nil
}
