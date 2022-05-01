package logic

import (
	"context"

	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ScreenDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewScreenDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenDetailLogic {
	return &ScreenDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据排片ID获取详情
func (l *ScreenDetailLogic) ScreenDetail(in *cinema.ScreenDetailRequest) (*cinema.ScreenDetailResponse, error) {

	resp := new(cinema.ScreenDetailResponse)

	screen, err := l.svcCtx.ScreenModel.FindOne(l.ctx, in.ScreenId)
	if err != nil {
		return &cinema.ScreenDetailResponse{}, errors.Wrap(err, "Cinema RPC:ScreenDetail [ScreenModel.FindOne] DbError")
	}

	copier.Copy(&resp, screen)

	return resp, nil
}
