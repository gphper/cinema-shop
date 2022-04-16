package logic

import (
	"context"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/usercenter/api/internal/svc"
	"cinema-shop/services/usercenter/api/internal/types"
	"cinema-shop/services/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {

	registResp, err := l.svcCtx.UserRpcClient.RegisterUser(l.ctx, &usercenter.RegisterRequest{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	})

	if err != nil {
		return nil, errorxx.NewCodeError(types.USER_REGISTER_ERR, "注册用户失败")
	}

	return &types.RegisterResp{
		Data: types.RegisterData{
			Id:           registResp.Id,
			Name:         registResp.Name,
			AccessExpire: registResp.Expire,
			AccessToken:  registResp.Token,
			RefreshToken: registResp.Reftoken,
		},
	}, nil
}
