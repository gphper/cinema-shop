// Code generated by goctl. DO NOT EDIT!
// Source: cinema.proto

package cinema

import (
	"context"

	"cinema-shop/services/cinema/rpc/pb/cinema"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FilmDatailRequest  = cinema.FilmDatailRequest
	FilmDetailInfo     = cinema.FilmDetailInfo
	FilmDetailResponse = cinema.FilmDetailResponse
	FilmListInfo       = cinema.FilmListInfo
	FilmListRequest    = cinema.FilmListRequest
	FilmListResponse   = cinema.FilmListResponse

	Cinema interface {
		// 影片列表
		List(ctx context.Context, in *FilmListRequest, opts ...grpc.CallOption) (*FilmListResponse, error)
		// 影片详情
		Detail(ctx context.Context, in *FilmDatailRequest, opts ...grpc.CallOption) (*FilmDetailResponse, error)
	}

	defaultCinema struct {
		cli zrpc.Client
	}
)

func NewCinema(cli zrpc.Client) Cinema {
	return &defaultCinema{
		cli: cli,
	}
}

// 影片列表
func (m *defaultCinema) List(ctx context.Context, in *FilmListRequest, opts ...grpc.CallOption) (*FilmListResponse, error) {
	client := cinema.NewCinemaClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

// 影片详情
func (m *defaultCinema) Detail(ctx context.Context, in *FilmDatailRequest, opts ...grpc.CallOption) (*FilmDetailResponse, error) {
	client := cinema.NewCinemaClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}