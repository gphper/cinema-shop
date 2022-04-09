package logic

import (
	"context"
	"time"

	"cinema-shop/common/jwtx"
	"cinema-shop/services/usercenter/rpc/internal/svc"
	"cinema-shop/services/usercenter/rpc/pb/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  刷新Token
func (l *RefreshLogic) Refresh(in *usercenter.RefreshRequest) (*usercenter.RefreshResponse, error) {

	userInfo, err := l.svcCtx.UserModel.FindOneByReftoken(l.ctx, in.Reftoken)
	if err != nil {
		return &usercenter.RefreshResponse{}, err
	}

	now := time.Now().Unix()
	jToken, err := jwtx.GetJwtToken(l.svcCtx.Config.AuthConfig.AccessSecret, now, l.svcCtx.Config.AuthConfig.AccessExpire, userInfo.Id)
	if err != nil {
		return &usercenter.RefreshResponse{}, err
	}

	return &usercenter.RefreshResponse{
		Id:       userInfo.Id,
		Name:     userInfo.Name,
		Token:    jToken,
		Expire:   l.svcCtx.Config.AuthConfig.AccessExpire,
		Reftoken: userInfo.Reftoken,
	}, nil
}
