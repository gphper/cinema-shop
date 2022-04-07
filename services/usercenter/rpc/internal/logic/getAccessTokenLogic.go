package logic

import (
	"context"

	"cinema-shop/services/usercenter/rpc/internal/svc"
	"cinema-shop/services/usercenter/rpc/pb/usercenter"

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

// func (l *GetAccessTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
// 	claims := make(jwt.MapClaims)
// 	claims["exp"] = iat + seconds
// 	claims["iat"] = iat
// 	claims["userId"] = userId
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	token.Claims = claims
// 	return token.SignedString([]byte(secretKey))
// }

//  根据邮箱获取用户秘钥和密码
func (l *GetAccessTokenLogic) GetAccessToken(in *usercenter.GetAccessTokenRequest) (*usercenter.GetAccessTokenResponse, error) {
	// todo: add your logic here and delete this line

	// userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	// if err != nil {
	// 	return &usercenter.GetAccessTokenResponse{}, err
	// }
	return &usercenter.GetAccessTokenResponse{}, nil
}
