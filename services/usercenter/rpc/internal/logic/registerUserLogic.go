package logic

import (
	"context"
	"errors"
	"time"

	"cinema-shop/common/jwtx"
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

	userObj, _ := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)

	if userObj != nil {
		return &usercenter.RegisterResponse{}, errors.New("邮箱信息已存在")
	}

	user := new(user.User)
	user.Email = in.Email
	user.Name = in.Name
	user.Salt = utilsx.RandString(10)
	user.Reftoken = uuid.NewString()
	user.Password = utilsx.Encryption(in.Password, user.Salt)
	l.svcCtx.UserModel.Insert(l.ctx, user)

	now := time.Now().Unix()
	jToken, err := jwtx.GetJwtToken(l.svcCtx.Config.AuthConfig.AccessSecret, now, l.svcCtx.Config.AuthConfig.AccessExpire, user.Id)
	if err != nil {
		return &usercenter.RegisterResponse{}, err
	}

	return &usercenter.RegisterResponse{
		Token:    jToken,
		Expire:   l.svcCtx.Config.AuthConfig.AccessExpire,
		Reftoken: user.Reftoken,
	}, nil
}
