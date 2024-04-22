// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: communication.proto

package __

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

// CentralServerClient is the client API for CentralServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CentralServerClient interface {
	SolicitarM(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type centralServerClient struct {
	cc grpc.ClientConnInterface
}

func NewCentralServerClient(cc grpc.ClientConnInterface) CentralServerClient {
	return &centralServerClient{cc}
}

func (c *centralServerClient) SolicitarM(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/communication.CentralServer/SolicitarM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentralServerServer is the server API for CentralServer service.
// All implementations must embed UnimplementedCentralServerServer
// for forward compatibility
type CentralServerServer interface {
	SolicitarM(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedCentralServerServer()
}

// UnimplementedCentralServerServer must be embedded to have forward compatible implementations.
type UnimplementedCentralServerServer struct {
}

func (UnimplementedCentralServerServer) SolicitarM(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarM not implemented")
}
func (UnimplementedCentralServerServer) mustEmbedUnimplementedCentralServerServer() {}

// UnsafeCentralServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CentralServerServer will
// result in compilation errors.
type UnsafeCentralServerServer interface {
	mustEmbedUnimplementedCentralServerServer()
}

func RegisterCentralServerServer(s grpc.ServiceRegistrar, srv CentralServerServer) {
	s.RegisterService(&CentralServer_ServiceDesc, srv)
}

func _CentralServer_SolicitarM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralServerServer).SolicitarM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.CentralServer/SolicitarM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralServerServer).SolicitarM(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CentralServer_ServiceDesc is the grpc.ServiceDesc for CentralServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CentralServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "communication.CentralServer",
	HandlerType: (*CentralServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolicitarM",
			Handler:    _CentralServer_SolicitarM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication.proto",
}