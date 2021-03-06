// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// RulesUIClient is the client API for RulesUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RulesUIClient interface {
	// ListEngineSpecs returns a list of Rules Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (RulesUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type rulesUIClient struct {
	cc grpc.ClientConnInterface
}

func NewRulesUIClient(cc grpc.ClientConnInterface) RulesUIClient {
	return &rulesUIClient{cc}
}

func (c *rulesUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (RulesUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RulesUI_ServiceDesc.Streams[0], "/v1.RulesUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &rulesUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RulesUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type rulesUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *rulesUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rulesUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.RulesUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesUIServer is the server API for RulesUI service.
// All implementations must embed UnimplementedRulesUIServer
// for forward compatibility
type RulesUIServer interface {
	// ListEngineSpecs returns a list of Rules Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, RulesUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedRulesUIServer()
}

// UnimplementedRulesUIServer must be embedded to have forward compatible implementations.
type UnimplementedRulesUIServer struct {
}

func (UnimplementedRulesUIServer) ListEngineSpecs(*ListEngineSpecsRequest, RulesUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedRulesUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedRulesUIServer) mustEmbedUnimplementedRulesUIServer() {}

// UnsafeRulesUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RulesUIServer will
// result in compilation errors.
type UnsafeRulesUIServer interface {
	mustEmbedUnimplementedRulesUIServer()
}

func RegisterRulesUIServer(s grpc.ServiceRegistrar, srv RulesUIServer) {
	s.RegisterService(&RulesUI_ServiceDesc, srv)
}

func _RulesUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RulesUIServer).ListEngineSpecs(m, &rulesUIListEngineSpecsServer{stream})
}

type RulesUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type rulesUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *rulesUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RulesUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.RulesUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RulesUI_ServiceDesc is the grpc.ServiceDesc for RulesUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RulesUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.RulesUI",
	HandlerType: (*RulesUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _RulesUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _RulesUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rules-ui.proto",
}
