package logic

import (
	"context"
	"errors"
	"time"

	"cinema-shop/common/jwtx"
	"cinema-shop/common/utilsx"
	"cinema-shop/services/usercenter/rpc/internal/svc"
	"cinema-shop/services/usercenter/rpc/pb/usercenter"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccessTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccessTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccessTokenLogic {
	return &GetAccessTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据邮箱获取用户秘钥和密码
func (l *GetAccessTokenLogic) GetAccessToken(in *usercenter.GetAccessTokenRequest) (*usercenter.GetAccessTokenResponse, error) {
	// todo: add your logic here and delete this line

	// userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	// if err != nil {
	// 	return &usercenter.GetAccessTokenResponse{}, err
	// }
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return &usercenter.GetAccessTokenResponse{}, err
	}

	password := utilsx.Encryption(in.Password, userInfo.Salt)
	if password != userInfo.Password {
		return &usercenter.GetAccessTokenResponse{}, errors.New("账号密码错误")
	}

	now := time.Now().Unix()
	jToken, err := jwtx.GetJwtToken(l.svcCtx.Config.AuthConfig.AccessSecret, now, l.svcCtx.Config.AuthConfig.AccessExpire, userInfo.Id)
	if err != nil {
		return &usercenter.GetAccessTokenResponse{}, err
	}

	//更新retoken
	userInfo.Reftoken = uuid.NewString()
	if err := l.svcCtx.UserModel.Update(l.ctx, userInfo); err != nil {
		return &usercenter.GetAccessTokenResponse{}, err
	}

	return &usercenter.GetAccessTokenResponse{
		Id:       userInfo.Id,
		Name:     userInfo.Name,
		Token:    jToken,
		Expire:   l.svcCtx.Config.AuthConfig.AccessExpire,
		Reftoken: userInfo.Reftoken,
	}, nil
}
