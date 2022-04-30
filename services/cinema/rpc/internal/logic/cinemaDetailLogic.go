package logic

import (
	"context"

	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CinemaDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCinemaDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CinemaDetailLogic {
	return &CinemaDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据影院ID获取详情
func (l *CinemaDetailLogic) CinemaDetail(in *cinema.CinemaDetailRequest) (*cinema.CinemaDetailResp, error) {

	var resp = new(cinema.CinemaDetailResp)

	cinemaInfo, err := l.svcCtx.CinemaModel.FindOne(l.ctx, in.CinemaId)
	if err != nil {
		return &cinema.CinemaDetailResp{}, errors.Wrap(err, "Cinema RPC:CinemaDetail Error")
	}

	copier.Copy(resp, &cinemaInfo)

	return resp, nil
}
