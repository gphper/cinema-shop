package logic

import (
	"context"

	filmm "cinema-shop/services/cinema/model/film"
	"cinema-shop/services/cinema/rpc/internal/svc"
	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 影片列表
func (l *ListLogic) List(in *cinema.FilmListRequest) (*cinema.FilmListResponse, error) {
	var (
		err   error
		data  filmm.PageLimitData
		count int
	)

	where := filmm.PageLimitWhere{
		Page:  int(in.Page),
		Limit: int(in.Limit),
	}

	err = l.svcCtx.FilmModel.PageLimit(l.ctx, where, &count, &data)
	if err != nil {
		return &cinema.FilmListResponse{}, err
	}

	var infos []*cinema.FilmListInfo
	if len(data) > 0 {
		infos = make([]*cinema.FilmListInfo, len(data))
		for k, v := range data {
			infos[k] = &cinema.FilmListInfo{
				FilmId:   int32(v.FilmId),
				FilmName: v.FilmName,
				CoverPic: v.CoverPic,
				Type:     int32(v.Tp),
			}
		}
	}

	return &cinema.FilmListResponse{
		Data:  infos,
		Count: int32(count),
		Page:  in.Page,
		Limit: in.Limit,
	}, nil
}
