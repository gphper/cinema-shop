package logic

import (
	"context"

	filmm "cinema-shop/services/film/model/film"
	"cinema-shop/services/film/rpc/internal/svc"
	"cinema-shop/services/film/rpc/pb/film"

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
func (l *ListLogic) List(in *film.FilmListRequest) (*film.FilmListResponse, error) {

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
		return &film.FilmListResponse{}, err
	}

	var infos []*film.FilmListInfo
	if len(data) > 0 {
		infos = make([]*film.FilmListInfo, len(data))
		for k, v := range data {
			infos[k] = &film.FilmListInfo{
				FilmId:   int32(v.FilmId),
				FilmName: v.FilmName,
				CoverPic: v.CoverPic,
				Type:     int32(v.Tp),
			}
		}
	}

	return &film.FilmListResponse{
		Data:  infos,
		Count: int32(count),
		Page:  in.Page,
		Limit: in.Limit,
	}, nil
}
