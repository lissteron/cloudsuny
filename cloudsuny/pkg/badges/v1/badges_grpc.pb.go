// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package badges

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BadgesClient is the client API for Badges service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BadgesClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
}

type badgesClient struct {
	cc grpc.ClientConnInterface
}

func NewBadgesClient(cc grpc.ClientConnInterface) BadgesClient {
	return &badgesClient{cc}
}

func (c *badgesClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, "/badges.v1.Badges/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *badgesClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/badges.v1.Badges/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BadgesServer is the server API for Badges service.
// All implementations must embed UnimplementedBadgesServer
// for forward compatibility
type BadgesServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	mustEmbedUnimplementedBadgesServer()
}

// UnimplementedBadgesServer must be embedded to have forward compatible implementations.
type UnimplementedBadgesServer struct {
}

func (UnimplementedBadgesServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBadgesServer) Update(context.Context, *UpdateReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBadgesServer) mustEmbedUnimplementedBadgesServer() {}

// UnsafeBadgesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BadgesServer will
// result in compilation errors.
type UnsafeBadgesServer interface {
	mustEmbedUnimplementedBadgesServer()
}

func RegisterBadgesServer(s *grpc.Server, srv BadgesServer) {
	s.RegisterService(&_Badges_serviceDesc, srv)
}

func _Badges_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badges.v1.Badges/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgesServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Badges_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BadgesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/badges.v1.Badges/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BadgesServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Badges_serviceDesc = grpc.ServiceDesc{
	ServiceName: "badges.v1.Badges",
	HandlerType: (*BadgesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Badges_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Badges_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "badges.proto",
}