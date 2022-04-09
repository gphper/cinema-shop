package logic

import (
	"context"

	"cinema-shop/services/usercenter/api/internal/svc"
	"cinema-shop/services/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Get_user_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet_user_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_user_infoLogic {
	return &Get_user_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_user_infoLogic) Get_user_info(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
