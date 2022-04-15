package logic

import (
	"context"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/usercenter/api/internal/svc"
	"cinema-shop/services/usercenter/api/internal/types"
	"cinema-shop/services/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {

	refRes, err := l.svcCtx.UserRpcClient.Refresh(l.ctx, &usercenter.RefreshRequest{
		Reftoken: req.RefToken,
	})

	if err != nil {
		return nil, errorxx.NewCustomError(types.USER_REFRESH_ERR, "获取新Token失败")
	}

	return &types.RefreshTokenResp{
		Data: types.RefreshTokenData{
			Id:           refRes.Id,
			Name:         refRes.Name,
			AccessToken:  refRes.Token,
			AccessExpire: refRes.Expire,
			RefreshToken: refRes.Reftoken,
		},
	}, nil
}
