// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BoxClient is the client API for Box service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoxClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Result, error)
	FormListQuestionType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FormListQuestionTypeResp, error)
	FormGet(ctx context.Context, in *FormGetReq, opts ...grpc.CallOption) (*FormGetResp, error)
	FormSave(ctx context.Context, in *FormSaveReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FormAnswerSave(ctx context.Context, in *FormSaveReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FormList(ctx context.Context, in *FormListReq, opts ...grpc.CallOption) (*FormListResp, error)
}

type boxClient struct {
	cc grpc.ClientConnInterface
}

func NewBoxClient(cc grpc.ClientConnInterface) BoxClient {
	return &boxClient{cc}
}

func (c *boxClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/golayout.v1.Box/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxClient) FormListQuestionType(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FormListQuestionTypeResp, error) {
	out := new(FormListQuestionTypeResp)
	err := c.cc.Invoke(ctx, "/golayout.v1.Box/FormListQuestionType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxClient) FormGet(ctx context.Context, in *FormGetReq, opts ...grpc.CallOption) (*FormGetResp, error) {
	out := new(FormGetResp)
	err := c.cc.Invoke(ctx, "/golayout.v1.Box/FormGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxClient) FormSave(ctx context.Context, in *FormSaveReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/golayout.v1.Box/FormSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxClient) FormAnswerSave(ctx context.Context, in *FormSaveReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/golayout.v1.Box/FormAnswerSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxClient) FormList(ctx context.Context, in *FormListReq, opts ...grpc.CallOption) (*FormListResp, error) {
	out := new(FormListResp)
	err := c.cc.Invoke(ctx, "/golayout.v1.Box/FormList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoxServer is the server API for Box service.
// All implementations must embed UnimplementedBoxServer
// for forward compatibility
type BoxServer interface {
	Ping(context.Context, *emptypb.Empty) (*Result, error)
	FormListQuestionType(context.Context, *emptypb.Empty) (*FormListQuestionTypeResp, error)
	FormGet(context.Context, *FormGetReq) (*FormGetResp, error)
	FormSave(context.Context, *FormSaveReq) (*emptypb.Empty, error)
	FormAnswerSave(context.Context, *FormSaveReq) (*emptypb.Empty, error)
	FormList(context.Context, *FormListReq) (*FormListResp, error)
	mustEmbedUnimplementedBoxServer()
}

// UnimplementedBoxServer must be embedded to have forward compatible implementations.
type UnimplementedBoxServer struct {
}

func (UnimplementedBoxServer) Ping(context.Context, *emptypb.Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBoxServer) FormListQuestionType(context.Context, *emptypb.Empty) (*FormListQuestionTypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormListQuestionType not implemented")
}
func (UnimplementedBoxServer) FormGet(context.Context, *FormGetReq) (*FormGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormGet not implemented")
}
func (UnimplementedBoxServer) FormSave(context.Context, *FormSaveReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormSave not implemented")
}
func (UnimplementedBoxServer) FormAnswerSave(context.Context, *FormSaveReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormAnswerSave not implemented")
}
func (UnimplementedBoxServer) FormList(context.Context, *FormListReq) (*FormListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormList not implemented")
}
func (UnimplementedBoxServer) mustEmbedUnimplementedBoxServer() {}

// UnsafeBoxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoxServer will
// result in compilation errors.
type UnsafeBoxServer interface {
	mustEmbedUnimplementedBoxServer()
}

func RegisterBoxServer(s grpc.ServiceRegistrar, srv BoxServer) {
	s.RegisterService(&Box_ServiceDesc, srv)
}

func _Box_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golayout.v1.Box/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Box_FormListQuestionType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServer).FormListQuestionType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golayout.v1.Box/FormListQuestionType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServer).FormListQuestionType(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Box_FormGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServer).FormGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golayout.v1.Box/FormGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServer).FormGet(ctx, req.(*FormGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Box_FormSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServer).FormSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golayout.v1.Box/FormSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServer).FormSave(ctx, req.(*FormSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Box_FormAnswerSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServer).FormAnswerSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golayout.v1.Box/FormAnswerSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServer).FormAnswerSave(ctx, req.(*FormSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Box_FormList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServer).FormList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golayout.v1.Box/FormList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServer).FormList(ctx, req.(*FormListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Box_ServiceDesc is the grpc.ServiceDesc for Box service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Box_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "golayout.v1.Box",
	HandlerType: (*BoxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Box_Ping_Handler,
		},
		{
			MethodName: "FormListQuestionType",
			Handler:    _Box_FormListQuestionType_Handler,
		},
		{
			MethodName: "FormGet",
			Handler:    _Box_FormGet_Handler,
		},
		{
			MethodName: "FormSave",
			Handler:    _Box_FormSave_Handler,
		},
		{
			MethodName: "FormAnswerSave",
			Handler:    _Box_FormAnswerSave_Handler,
		},
		{
			MethodName: "FormList",
			Handler:    _Box_FormList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/box.proto",
}
