package logic

import (
	"context"

	"cinema-shop/services/usercenter/api/internal/svc"
	"cinema-shop/services/usercenter/api/internal/types"
	"cinema-shop/services/usercenter/rpc/pb/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	userResp, err := l.svcCtx.UserRpcClient.GetAccessToken(l.ctx, &usercenter.GetAccessTokenRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return
	}

	return &types.LoginResp{
		Id:           userResp.Id,
		Name:         userResp.Name,
		AccessToken:  userResp.Token,
		AccessExpire: userResp.Expire,
		RefreshToken: userResp.Reftoken,
	}, nil

}
