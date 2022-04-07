package logic

import (
	"context"

	"cinema-shop/common/utilsx"
	"cinema-shop/services/usercenter/model/user"
	"cinema-shop/services/usercenter/rpc/internal/svc"
	"cinema-shop/services/usercenter/rpc/pb/usercenter"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册
func (l *RegisterUserLogic) RegisterUser(in *usercenter.RegisterRequest) (*usercenter.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	user := new(user.User)
	user.Email = in.Email
	user.Name = in.Name
	user.Salt = utilsx.RandString(10)
	user.Reftoken = uuid.NewString()

	l.svcCtx.UserModel.Insert(l.ctx, user)

	return &usercenter.RegisterResponse{}, nil
}
