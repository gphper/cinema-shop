package logic

import (
	"context"
	"errors"
	"time"

	"cinema-shop/services/user/api/internal/svc"
	"cinema-shop/services/user/api/internal/types"
	"cinema-shop/services/user/rpc/user"

	"github.com/golang-jwt/jwt/v4"
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

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	userResp, err := l.svcCtx.UserRpcClient.GetUserByEmail(l.ctx, &user.GetUserByEmailRequest{
		Email: req.Email,
	})

	if err != nil {
		return
	}

	if req.Password != userResp.Password {
		return nil, errors.New("用户密码不正确")
	}

	now := time.Now().Unix()
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userResp.Id)

	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Id:           userResp.Id,
		Name:         userResp.Name,
		AccessToken:  jwtToken,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
