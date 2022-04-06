package logic

import (
	"context"

	"cinema-shop/services/user/api/internal/svc"
	"cinema-shop/services/user/api/internal/types"
	"cinema-shop/services/user/rpc/user"

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
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserRpcClient.GetUserByEmail(l.ctx, &user.GetUserByEmailRequest{
		Email: req.Email,
	})
	if err != nil {
		return
	}

	logx.Info(userResp)

	return
}
