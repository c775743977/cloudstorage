// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: login.proto

package pbfile

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

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Regist(ctx context.Context, in *RegistRequest, opts ...grpc.CallOption) (*RegistResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	GetSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Cookie, error)
	RegistVIP(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pbfile.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Regist(ctx context.Context, in *RegistRequest, opts ...grpc.CallOption) (*RegistResponse, error) {
	out := new(RegistResponse)
	err := c.cc.Invoke(ctx, "/pbfile.LoginService/Regist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/pbfile.LoginService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) GetSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pbfile.LoginService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Logout(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, "/pbfile.LoginService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) RegistVIP(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/pbfile.LoginService/RegistVIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Regist(context.Context, *RegistRequest) (*RegistResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	GetSession(context.Context, *Cookie) (*LoginResponse, error)
	Logout(context.Context, *Cookie) (*Cookie, error)
	RegistVIP(context.Context, *Cookie) (*VerifyResponse, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServiceServer) Regist(context.Context, *RegistRequest) (*RegistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Regist not implemented")
}
func (UnimplementedLoginServiceServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedLoginServiceServer) GetSession(context.Context, *Cookie) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedLoginServiceServer) Logout(context.Context, *Cookie) (*Cookie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedLoginServiceServer) RegistVIP(context.Context, *Cookie) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistVIP not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Regist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Regist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.LoginService/Regist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Regist(ctx, req.(*RegistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.LoginService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.LoginService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).GetSession(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.LoginService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Logout(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_RegistVIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).RegistVIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbfile.LoginService/RegistVIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).RegistVIP(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbfile.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
		{
			MethodName: "Regist",
			Handler:    _LoginService_Regist_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _LoginService_Verify_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _LoginService_GetSession_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _LoginService_Logout_Handler,
		},
		{
			MethodName: "RegistVIP",
			Handler:    _LoginService_RegistVIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
