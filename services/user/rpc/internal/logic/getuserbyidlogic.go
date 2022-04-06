package logic

import (
	"context"

	"cinema-shop/services/user/rpc/internal/svc"
	"cinema-shop/services/user/rpc/pb/user"

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
func (l *GetUserByIDLogic) GetUserByID(in *user.GetUserByIDRequest) (*user.GetUserByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserByIdResponse{}, nil
}
