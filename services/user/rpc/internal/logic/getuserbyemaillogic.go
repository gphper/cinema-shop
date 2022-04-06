package logic

import (
	"context"

	"cinema-shop/services/user/rpc/internal/svc"
	"cinema-shop/services/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据邮箱获取用户秘钥和密码
func (l *GetUserByEmailLogic) GetUserByEmail(in *user.GetUserByEmailRequest) (*user.GetUserByEmailResponse, error) {

	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return &user.GetUserByEmailResponse{}, err
	}

	return &user.GetUserByEmailResponse{
		Password: userInfo.Password,
		Salt:     userInfo.Salt,
		Id:       userInfo.Id,
	}, nil
}
