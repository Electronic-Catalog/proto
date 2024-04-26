// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: cmng/cmng.proto

package cmng

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

// CardManagementServiceClient is the client API for CardManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardManagementServiceClient interface {
	SetCard(ctx context.Context, in *SetCardRequest, opts ...grpc.CallOption) (*SetCardResponse, error)
	GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error)
}

type cardManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardManagementServiceClient(cc grpc.ClientConnInterface) CardManagementServiceClient {
	return &cardManagementServiceClient{cc}
}

func (c *cardManagementServiceClient) SetCard(ctx context.Context, in *SetCardRequest, opts ...grpc.CallOption) (*SetCardResponse, error) {
	out := new(SetCardResponse)
	err := c.cc.Invoke(ctx, "/cmng.CardManagementService/setCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardManagementServiceClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error) {
	out := new(GetCardResponse)
	err := c.cc.Invoke(ctx, "/cmng.CardManagementService/getCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardManagementServiceServer is the server API for CardManagementService service.
// All implementations should embed UnimplementedCardManagementServiceServer
// for forward compatibility
type CardManagementServiceServer interface {
	SetCard(context.Context, *SetCardRequest) (*SetCardResponse, error)
	GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error)
}

// UnimplementedCardManagementServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCardManagementServiceServer struct {
}

func (UnimplementedCardManagementServiceServer) SetCard(context.Context, *SetCardRequest) (*SetCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCard not implemented")
}
func (UnimplementedCardManagementServiceServer) GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}

// UnsafeCardManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardManagementServiceServer will
// result in compilation errors.
type UnsafeCardManagementServiceServer interface {
	mustEmbedUnimplementedCardManagementServiceServer()
}

func RegisterCardManagementServiceServer(s grpc.ServiceRegistrar, srv CardManagementServiceServer) {
	s.RegisterService(&CardManagementService_ServiceDesc, srv)
}

func _CardManagementService_SetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardManagementServiceServer).SetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cmng.CardManagementService/setCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardManagementServiceServer).SetCard(ctx, req.(*SetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardManagementService_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardManagementServiceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cmng.CardManagementService/getCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardManagementServiceServer).GetCard(ctx, req.(*GetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardManagementService_ServiceDesc is the grpc.ServiceDesc for CardManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cmng.CardManagementService",
	HandlerType: (*CardManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "setCard",
			Handler:    _CardManagementService_SetCard_Handler,
		},
		{
			MethodName: "getCard",
			Handler:    _CardManagementService_GetCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmng/cmng.proto",
}
