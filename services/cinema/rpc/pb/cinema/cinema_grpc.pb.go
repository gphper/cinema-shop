// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: services/cinema/rpc/pb/cinema/cinema.proto

package cinema

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CinemaClient is the client API for Cinema service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CinemaClient interface {
	//影片列表
	FilmList(ctx context.Context, in *FilmListRequest, opts ...grpc.CallOption) (*FilmListResponse, error)
	//影片详情
	FilmDetail(ctx context.Context, in *FilmDatailRequest, opts ...grpc.CallOption) (*FilmDetailResponse, error)
	//根据地理位置获取影院信息
	CinemaInfo(ctx context.Context, in *CinemaInfoRequest, opts ...grpc.CallOption) (*CinemaInfoResponse, error)
	//根据影片和日期和影院ID获取排片信息
	ScreenCinemaInfo(ctx context.Context, in *ScreenCinemaInfoRequest, opts ...grpc.CallOption) (*ScreenCinemaInfoResponse, error)
	//根据影院ID获取详情
	CinemaDetail(ctx context.Context, in *CinemaDetailRequest, opts ...grpc.CallOption) (*CinemaDetailResp, error)
	//根据日期、影院ID获取排片电影
	ScreenFilmId(ctx context.Context, in *ScreenFilmIdRequest, opts ...grpc.CallOption) (*ScreenFilmIdResponse, error)
	//根据影片ID获取全部影片信息
	FilmAll(ctx context.Context, in *FilmAllRequest, opts ...grpc.CallOption) (*FilmAllResponse, error)
}

type cinemaClient struct {
	cc grpc.ClientConnInterface
}

func NewCinemaClient(cc grpc.ClientConnInterface) CinemaClient {
	return &cinemaClient{cc}
}

func (c *cinemaClient) FilmList(ctx context.Context, in *FilmListRequest, opts ...grpc.CallOption) (*FilmListResponse, error) {
	out := new(FilmListResponse)
	err := c.cc.Invoke(ctx, "/cinema.Cinema/FilmList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaClient) FilmDetail(ctx context.Context, in *FilmDatailRequest, opts ...grpc.CallOption) (*FilmDetailResponse, error) {
	out := new(FilmDetailResponse)
	err := c.cc.Invoke(ctx, "/cinema.Cinema/FilmDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaClient) CinemaInfo(ctx context.Context, in *CinemaInfoRequest, opts ...grpc.CallOption) (*CinemaInfoResponse, error) {
	out := new(CinemaInfoResponse)
	err := c.cc.Invoke(ctx, "/cinema.Cinema/CinemaInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaClient) ScreenCinemaInfo(ctx context.Context, in *ScreenCinemaInfoRequest, opts ...grpc.CallOption) (*ScreenCinemaInfoResponse, error) {
	out := new(ScreenCinemaInfoResponse)
	err := c.cc.Invoke(ctx, "/cinema.Cinema/ScreenCinemaInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaClient) CinemaDetail(ctx context.Context, in *CinemaDetailRequest, opts ...grpc.CallOption) (*CinemaDetailResp, error) {
	out := new(CinemaDetailResp)
	err := c.cc.Invoke(ctx, "/cinema.Cinema/CinemaDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaClient) ScreenFilmId(ctx context.Context, in *ScreenFilmIdRequest, opts ...grpc.CallOption) (*ScreenFilmIdResponse, error) {
	out := new(ScreenFilmIdResponse)
	err := c.cc.Invoke(ctx, "/cinema.Cinema/ScreenFilmId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaClient) FilmAll(ctx context.Context, in *FilmAllRequest, opts ...grpc.CallOption) (*FilmAllResponse, error) {
	out := new(FilmAllResponse)
	err := c.cc.Invoke(ctx, "/cinema.Cinema/FilmAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CinemaServer is the server API for Cinema service.
// All implementations must embed UnimplementedCinemaServer
// for forward compatibility
type CinemaServer interface {
	//影片列表
	FilmList(context.Context, *FilmListRequest) (*FilmListResponse, error)
	//影片详情
	FilmDetail(context.Context, *FilmDatailRequest) (*FilmDetailResponse, error)
	//根据地理位置获取影院信息
	CinemaInfo(context.Context, *CinemaInfoRequest) (*CinemaInfoResponse, error)
	//根据影片和日期和影院ID获取排片信息
	ScreenCinemaInfo(context.Context, *ScreenCinemaInfoRequest) (*ScreenCinemaInfoResponse, error)
	//根据影院ID获取详情
	CinemaDetail(context.Context, *CinemaDetailRequest) (*CinemaDetailResp, error)
	//根据日期、影院ID获取排片电影
	ScreenFilmId(context.Context, *ScreenFilmIdRequest) (*ScreenFilmIdResponse, error)
	//根据影片ID获取全部影片信息
	FilmAll(context.Context, *FilmAllRequest) (*FilmAllResponse, error)
	mustEmbedUnimplementedCinemaServer()
}

// UnimplementedCinemaServer must be embedded to have forward compatible implementations.
type UnimplementedCinemaServer struct {
}

func (UnimplementedCinemaServer) FilmList(context.Context, *FilmListRequest) (*FilmListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilmList not implemented")
}
func (UnimplementedCinemaServer) FilmDetail(context.Context, *FilmDatailRequest) (*FilmDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilmDetail not implemented")
}
func (UnimplementedCinemaServer) CinemaInfo(context.Context, *CinemaInfoRequest) (*CinemaInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CinemaInfo not implemented")
}
func (UnimplementedCinemaServer) ScreenCinemaInfo(context.Context, *ScreenCinemaInfoRequest) (*ScreenCinemaInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScreenCinemaInfo not implemented")
}
func (UnimplementedCinemaServer) CinemaDetail(context.Context, *CinemaDetailRequest) (*CinemaDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CinemaDetail not implemented")
}
func (UnimplementedCinemaServer) ScreenFilmId(context.Context, *ScreenFilmIdRequest) (*ScreenFilmIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScreenFilmId not implemented")
}
func (UnimplementedCinemaServer) FilmAll(context.Context, *FilmAllRequest) (*FilmAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilmAll not implemented")
}
func (UnimplementedCinemaServer) mustEmbedUnimplementedCinemaServer() {}

// UnsafeCinemaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CinemaServer will
// result in compilation errors.
type UnsafeCinemaServer interface {
	mustEmbedUnimplementedCinemaServer()
}

func RegisterCinemaServer(s grpc.ServiceRegistrar, srv CinemaServer) {
	s.RegisterService(&Cinema_ServiceDesc, srv)
}

func _Cinema_FilmList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServer).FilmList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.Cinema/FilmList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServer).FilmList(ctx, req.(*FilmListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cinema_FilmDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmDatailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServer).FilmDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.Cinema/FilmDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServer).FilmDetail(ctx, req.(*FilmDatailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cinema_CinemaInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CinemaInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServer).CinemaInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.Cinema/CinemaInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServer).CinemaInfo(ctx, req.(*CinemaInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cinema_ScreenCinemaInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScreenCinemaInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServer).ScreenCinemaInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.Cinema/ScreenCinemaInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServer).ScreenCinemaInfo(ctx, req.(*ScreenCinemaInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cinema_CinemaDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CinemaDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServer).CinemaDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.Cinema/CinemaDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServer).CinemaDetail(ctx, req.(*CinemaDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cinema_ScreenFilmId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScreenFilmIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServer).ScreenFilmId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.Cinema/ScreenFilmId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServer).ScreenFilmId(ctx, req.(*ScreenFilmIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cinema_FilmAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilmAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServer).FilmAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.Cinema/FilmAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServer).FilmAll(ctx, req.(*FilmAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cinema_ServiceDesc is the grpc.ServiceDesc for Cinema service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cinema_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cinema.Cinema",
	HandlerType: (*CinemaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FilmList",
			Handler:    _Cinema_FilmList_Handler,
		},
		{
			MethodName: "FilmDetail",
			Handler:    _Cinema_FilmDetail_Handler,
		},
		{
			MethodName: "CinemaInfo",
			Handler:    _Cinema_CinemaInfo_Handler,
		},
		{
			MethodName: "ScreenCinemaInfo",
			Handler:    _Cinema_ScreenCinemaInfo_Handler,
		},
		{
			MethodName: "CinemaDetail",
			Handler:    _Cinema_CinemaDetail_Handler,
		},
		{
			MethodName: "ScreenFilmId",
			Handler:    _Cinema_ScreenFilmId_Handler,
		},
		{
			MethodName: "FilmAll",
			Handler:    _Cinema_FilmAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/cinema/rpc/pb/cinema/cinema.proto",
}
