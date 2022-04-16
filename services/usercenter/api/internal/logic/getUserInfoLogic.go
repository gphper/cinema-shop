package logic

import (
	"context"
	"encoding/json"

	"cinema-shop/common/errorxx"
	"cinema-shop/services/usercenter/api/internal/svc"
	"cinema-shop/services/usercenter/api/internal/types"
	"cinema-shop/services/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {

	id, err := l.ctx.Value("userId").(json.Number).Int64()

	if err != nil {
		return
	}

	userInfo, err := l.svcCtx.UserRpcClient.GetUserByID(l.ctx, &usercenter.GetUserByIDRequest{
		Id: id,
	})

	if err != nil {
		return nil, errorxx.NewCodeError(types.USER_INFO_ERR, "获取用户信息失败")
	}

	resp = &types.GetUserInfoResp{
		Data: types.GetUserInfoData{
			Id:    userInfo.Info.Id,
			Name:  userInfo.Info.Name,
			Email: userInfo.Info.Email,
		},
	}

	return
}
