// we will need to run protoc --go_out=. --go_opt=paths=source_relative \
//--go-grpc_out=. --go-grpc_opt=paths=source_relative \
//Proto/mail.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/mail.proto

package proto

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

const (
	MailingListService_CreateEmail_FullMethodName   = "/proto.MailingListService/CreateEmail"
	MailingListService_GetEmail_FullMethodName      = "/proto.MailingListService/GetEmail"
	MailingListService_UpdateEmail_FullMethodName   = "/proto.MailingListService/UpdateEmail"
	MailingListService_DeleteEmail_FullMethodName   = "/proto.MailingListService/DeleteEmail"
	MailingListService_GetEmailBatch_FullMethodName = "/proto.MailingListService/GetEmailBatch"
)

// MailingListServiceClient is the client API for MailingListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailingListServiceClient interface {
	CreateEmail(ctx context.Context, in *CreateEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
	GetEmail(ctx context.Context, in *GetEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
	UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
	DeleteEmail(ctx context.Context, in *DeleteEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
	GetEmailBatch(ctx context.Context, in *GetEmailBatchRequest, opts ...grpc.CallOption) (*GetEmailBatchResponse, error)
}

type mailingListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailingListServiceClient(cc grpc.ClientConnInterface) MailingListServiceClient {
	return &mailingListServiceClient{cc}
}

func (c *mailingListServiceClient) CreateEmail(ctx context.Context, in *CreateEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, MailingListService_CreateEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailingListServiceClient) GetEmail(ctx context.Context, in *GetEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, MailingListService_GetEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailingListServiceClient) UpdateEmail(ctx context.Context, in *UpdateEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, MailingListService_UpdateEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailingListServiceClient) DeleteEmail(ctx context.Context, in *DeleteEmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, MailingListService_DeleteEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailingListServiceClient) GetEmailBatch(ctx context.Context, in *GetEmailBatchRequest, opts ...grpc.CallOption) (*GetEmailBatchResponse, error) {
	out := new(GetEmailBatchResponse)
	err := c.cc.Invoke(ctx, MailingListService_GetEmailBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailingListServiceServer is the server API for MailingListService service.
// All implementations must embed UnimplementedMailingListServiceServer
// for forward compatibility
type MailingListServiceServer interface {
	CreateEmail(context.Context, *CreateEmailRequest) (*EmailResponse, error)
	GetEmail(context.Context, *GetEmailRequest) (*EmailResponse, error)
	UpdateEmail(context.Context, *UpdateEmailRequest) (*EmailResponse, error)
	DeleteEmail(context.Context, *DeleteEmailRequest) (*EmailResponse, error)
	GetEmailBatch(context.Context, *GetEmailBatchRequest) (*GetEmailBatchResponse, error)
	mustEmbedUnimplementedMailingListServiceServer()
}

// UnimplementedMailingListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMailingListServiceServer struct {
}

func (UnimplementedMailingListServiceServer) CreateEmail(context.Context, *CreateEmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmail not implemented")
}
func (UnimplementedMailingListServiceServer) GetEmail(context.Context, *GetEmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmail not implemented")
}
func (UnimplementedMailingListServiceServer) UpdateEmail(context.Context, *UpdateEmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmail not implemented")
}
func (UnimplementedMailingListServiceServer) DeleteEmail(context.Context, *DeleteEmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmail not implemented")
}
func (UnimplementedMailingListServiceServer) GetEmailBatch(context.Context, *GetEmailBatchRequest) (*GetEmailBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailBatch not implemented")
}
func (UnimplementedMailingListServiceServer) mustEmbedUnimplementedMailingListServiceServer() {}

// UnsafeMailingListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailingListServiceServer will
// result in compilation errors.
type UnsafeMailingListServiceServer interface {
	mustEmbedUnimplementedMailingListServiceServer()
}

func RegisterMailingListServiceServer(s grpc.ServiceRegistrar, srv MailingListServiceServer) {
	s.RegisterService(&MailingListService_ServiceDesc, srv)
}

func _MailingListService_CreateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServiceServer).CreateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailingListService_CreateEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServiceServer).CreateEmail(ctx, req.(*CreateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailingListService_GetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServiceServer).GetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailingListService_GetEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServiceServer).GetEmail(ctx, req.(*GetEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailingListService_UpdateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServiceServer).UpdateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailingListService_UpdateEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServiceServer).UpdateEmail(ctx, req.(*UpdateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailingListService_DeleteEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServiceServer).DeleteEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailingListService_DeleteEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServiceServer).DeleteEmail(ctx, req.(*DeleteEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailingListService_GetEmailBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEmailBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailingListServiceServer).GetEmailBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailingListService_GetEmailBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailingListServiceServer).GetEmailBatch(ctx, req.(*GetEmailBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailingListService_ServiceDesc is the grpc.ServiceDesc for MailingListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailingListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MailingListService",
	HandlerType: (*MailingListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmail",
			Handler:    _MailingListService_CreateEmail_Handler,
		},
		{
			MethodName: "GetEmail",
			Handler:    _MailingListService_GetEmail_Handler,
		},
		{
			MethodName: "UpdateEmail",
			Handler:    _MailingListService_UpdateEmail_Handler,
		},
		{
			MethodName: "DeleteEmail",
			Handler:    _MailingListService_DeleteEmail_Handler,
		},
		{
			MethodName: "GetEmailBatch",
			Handler:    _MailingListService_GetEmailBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mail.proto",
}
