package logic

import (
	"context"

	"cinema-shop/services/usercenter/rpc/internal/svc"
	"cinema-shop/services/usercenter/rpc/pb/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIDLogic {
	return &GetUserByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据ID获取信息
func (l *GetUserByIDLogic) GetUserByID(in *usercenter.GetUserByIDRequest) (*usercenter.GetUserByIdResponse, error) {

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return &usercenter.GetUserByIdResponse{}, err
	}

	return &usercenter.GetUserByIdResponse{
		Info: &usercenter.UserInfo{
			Id:    userInfo.Id,
			Name:  userInfo.Name,
			Email: userInfo.Email,
		},
	}, nil
}
