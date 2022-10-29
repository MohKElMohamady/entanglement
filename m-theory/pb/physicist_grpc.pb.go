// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/physicist.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PhysicistsInfoClient is the client API for PhysicistsInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhysicistsInfoClient interface {
	GetPhysicistById(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Physicist, error)
	AddPhysicist(ctx context.Context, in *Physicist, opts ...grpc.CallOption) (*Physicist, error)
	GetAllPhysicist(ctx context.Context, in *AllPhysicistsRequest, opts ...grpc.CallOption) (*AllPhysicists, error)
	GetPhysicistsByCountryOfBirth(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (PhysicistsInfo_GetPhysicistsByCountryOfBirthClient, error)
}

type physicistsInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewPhysicistsInfoClient(cc grpc.ClientConnInterface) PhysicistsInfoClient {
	return &physicistsInfoClient{cc}
}

func (c *physicistsInfoClient) GetPhysicistById(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Physicist, error) {
	out := new(Physicist)
	err := c.cc.Invoke(ctx, "/physicist_info.PhysicistsInfo/getPhysicistById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *physicistsInfoClient) AddPhysicist(ctx context.Context, in *Physicist, opts ...grpc.CallOption) (*Physicist, error) {
	out := new(Physicist)
	err := c.cc.Invoke(ctx, "/physicist_info.PhysicistsInfo/addPhysicist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *physicistsInfoClient) GetAllPhysicist(ctx context.Context, in *AllPhysicistsRequest, opts ...grpc.CallOption) (*AllPhysicists, error) {
	out := new(AllPhysicists)
	err := c.cc.Invoke(ctx, "/physicist_info.PhysicistsInfo/getAllPhysicist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *physicistsInfoClient) GetPhysicistsByCountryOfBirth(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (PhysicistsInfo_GetPhysicistsByCountryOfBirthClient, error) {
	stream, err := c.cc.NewStream(ctx, &PhysicistsInfo_ServiceDesc.Streams[0], "/physicist_info.PhysicistsInfo/getPhysicistsByCountryOfBirth", opts...)
	if err != nil {
		return nil, err
	}
	x := &physicistsInfoGetPhysicistsByCountryOfBirthClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PhysicistsInfo_GetPhysicistsByCountryOfBirthClient interface {
	Recv() (*Physicist, error)
	grpc.ClientStream
}

type physicistsInfoGetPhysicistsByCountryOfBirthClient struct {
	grpc.ClientStream
}

func (x *physicistsInfoGetPhysicistsByCountryOfBirthClient) Recv() (*Physicist, error) {
	m := new(Physicist)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PhysicistsInfoServer is the server API for PhysicistsInfo service.
// All implementations must embed UnimplementedPhysicistsInfoServer
// for forward compatibility
type PhysicistsInfoServer interface {
	GetPhysicistById(context.Context, *UUID) (*Physicist, error)
	AddPhysicist(context.Context, *Physicist) (*Physicist, error)
	GetAllPhysicist(context.Context, *AllPhysicistsRequest) (*AllPhysicists, error)
	GetPhysicistsByCountryOfBirth(*wrapperspb.StringValue, PhysicistsInfo_GetPhysicistsByCountryOfBirthServer) error
	mustEmbedUnimplementedPhysicistsInfoServer()
}

// UnimplementedPhysicistsInfoServer must be embedded to have forward compatible implementations.
type UnimplementedPhysicistsInfoServer struct {
}

func (UnimplementedPhysicistsInfoServer) GetPhysicistById(context.Context, *UUID) (*Physicist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhysicistById not implemented")
}
func (UnimplementedPhysicistsInfoServer) AddPhysicist(context.Context, *Physicist) (*Physicist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPhysicist not implemented")
}
func (UnimplementedPhysicistsInfoServer) GetAllPhysicist(context.Context, *AllPhysicistsRequest) (*AllPhysicists, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPhysicist not implemented")
}
func (UnimplementedPhysicistsInfoServer) GetPhysicistsByCountryOfBirth(*wrapperspb.StringValue, PhysicistsInfo_GetPhysicistsByCountryOfBirthServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPhysicistsByCountryOfBirth not implemented")
}
func (UnimplementedPhysicistsInfoServer) mustEmbedUnimplementedPhysicistsInfoServer() {}

// UnsafePhysicistsInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhysicistsInfoServer will
// result in compilation errors.
type UnsafePhysicistsInfoServer interface {
	mustEmbedUnimplementedPhysicistsInfoServer()
}

func RegisterPhysicistsInfoServer(s grpc.ServiceRegistrar, srv PhysicistsInfoServer) {
	s.RegisterService(&PhysicistsInfo_ServiceDesc, srv)
}

func _PhysicistsInfo_GetPhysicistById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhysicistsInfoServer).GetPhysicistById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/physicist_info.PhysicistsInfo/getPhysicistById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhysicistsInfoServer).GetPhysicistById(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhysicistsInfo_AddPhysicist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Physicist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhysicistsInfoServer).AddPhysicist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/physicist_info.PhysicistsInfo/addPhysicist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhysicistsInfoServer).AddPhysicist(ctx, req.(*Physicist))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhysicistsInfo_GetAllPhysicist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllPhysicistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhysicistsInfoServer).GetAllPhysicist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/physicist_info.PhysicistsInfo/getAllPhysicist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhysicistsInfoServer).GetAllPhysicist(ctx, req.(*AllPhysicistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhysicistsInfo_GetPhysicistsByCountryOfBirth_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PhysicistsInfoServer).GetPhysicistsByCountryOfBirth(m, &physicistsInfoGetPhysicistsByCountryOfBirthServer{stream})
}

type PhysicistsInfo_GetPhysicistsByCountryOfBirthServer interface {
	Send(*Physicist) error
	grpc.ServerStream
}

type physicistsInfoGetPhysicistsByCountryOfBirthServer struct {
	grpc.ServerStream
}

func (x *physicistsInfoGetPhysicistsByCountryOfBirthServer) Send(m *Physicist) error {
	return x.ServerStream.SendMsg(m)
}

// PhysicistsInfo_ServiceDesc is the grpc.ServiceDesc for PhysicistsInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhysicistsInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "physicist_info.PhysicistsInfo",
	HandlerType: (*PhysicistsInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getPhysicistById",
			Handler:    _PhysicistsInfo_GetPhysicistById_Handler,
		},
		{
			MethodName: "addPhysicist",
			Handler:    _PhysicistsInfo_AddPhysicist_Handler,
		},
		{
			MethodName: "getAllPhysicist",
			Handler:    _PhysicistsInfo_GetAllPhysicist_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getPhysicistsByCountryOfBirth",
			Handler:       _PhysicistsInfo_GetPhysicistsByCountryOfBirth_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/physicist.proto",
}
