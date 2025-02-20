// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: http/cas/v1/sysuser.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Sysuser_CreateSysuser_FullMethodName       = "/api.http.cas.v1.Sysuser/CreateSysuser"
	Sysuser_UpdateSysuser_FullMethodName       = "/api.http.cas.v1.Sysuser/UpdateSysuser"
	Sysuser_DeleteSysuser_FullMethodName       = "/api.http.cas.v1.Sysuser/DeleteSysuser"
	Sysuser_GetSysuser_FullMethodName          = "/api.http.cas.v1.Sysuser/GetSysuser"
	Sysuser_ListSysuser_FullMethodName         = "/api.http.cas.v1.Sysuser/ListSysuser"
	Sysuser_GetCaptcha_FullMethodName          = "/api.http.cas.v1.Sysuser/GetCaptcha"
	Sysuser_Login_FullMethodName               = "/api.http.cas.v1.Sysuser/Login"
	Sysuser_Logout_FullMethodName              = "/api.http.cas.v1.Sysuser/Logout"
	Sysuser_Auth_FullMethodName                = "/api.http.cas.v1.Sysuser/Auth"
	Sysuser_ChangeStatus_FullMethodName        = "/api.http.cas.v1.Sysuser/ChangeStatus"
	Sysuser_UpdatePassword_FullMethodName      = "/api.http.cas.v1.Sysuser/UpdatePassword"
	Sysuser_GetPostInit_FullMethodName         = "/api.http.cas.v1.Sysuser/GetPostInit"
	Sysuser_GetUserRolePost_FullMethodName     = "/api.http.cas.v1.Sysuser/GetUserRolePost"
	Sysuser_GetUserGoogleSecret_FullMethodName = "/api.http.cas.v1.Sysuser/GetUserGoogleSecret"
)

// SysuserClient is the client API for Sysuser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 用户管理
type SysuserClient interface {
	// 创建用户
	CreateSysuser(ctx context.Context, in *CreateSysuserRequest, opts ...grpc.CallOption) (*CreateSysuserReply, error)
	// 更新用户
	UpdateSysuser(ctx context.Context, in *UpdateSysuserRequest, opts ...grpc.CallOption) (*UpdateSysuserReply, error)
	// 删除用户
	DeleteSysuser(ctx context.Context, in *DeleteSysuserRequest, opts ...grpc.CallOption) (*DeleteSysuserReply, error)
	// 获取用户
	GetSysuser(ctx context.Context, in *GetSysuserRequest, opts ...grpc.CallOption) (*GetSysuserReply, error)
	// 用户列表
	ListSysuser(ctx context.Context, in *ListSysuserRequest, opts ...grpc.CallOption) (*ListSysuserReply, error)
	// 获取验证码
	GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaReply, error)
	// 登入
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 登出
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	// 获取用户权限
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	// 更新用户状态
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatusReply, error)
	// 更新密码
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordReply, error)
	// 获取岗位
	GetPostInit(ctx context.Context, in *GetPostInitRequest, opts ...grpc.CallOption) (*GetPostInitReply, error)
	// 获取RoPo
	GetUserRolePost(ctx context.Context, in *GetUserRolePostRequest, opts ...grpc.CallOption) (*GetUserRolePostReply, error)
	// 生成密钥和二维码
	GetUserGoogleSecret(ctx context.Context, in *GetUserGoogleSecretRequest, opts ...grpc.CallOption) (*GetUserGoogleSecretReply, error)
}

type sysuserClient struct {
	cc grpc.ClientConnInterface
}

func NewSysuserClient(cc grpc.ClientConnInterface) SysuserClient {
	return &sysuserClient{cc}
}

func (c *sysuserClient) CreateSysuser(ctx context.Context, in *CreateSysuserRequest, opts ...grpc.CallOption) (*CreateSysuserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSysuserReply)
	err := c.cc.Invoke(ctx, Sysuser_CreateSysuser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) UpdateSysuser(ctx context.Context, in *UpdateSysuserRequest, opts ...grpc.CallOption) (*UpdateSysuserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSysuserReply)
	err := c.cc.Invoke(ctx, Sysuser_UpdateSysuser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) DeleteSysuser(ctx context.Context, in *DeleteSysuserRequest, opts ...grpc.CallOption) (*DeleteSysuserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSysuserReply)
	err := c.cc.Invoke(ctx, Sysuser_DeleteSysuser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) GetSysuser(ctx context.Context, in *GetSysuserRequest, opts ...grpc.CallOption) (*GetSysuserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSysuserReply)
	err := c.cc.Invoke(ctx, Sysuser_GetSysuser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) ListSysuser(ctx context.Context, in *ListSysuserRequest, opts ...grpc.CallOption) (*ListSysuserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSysuserReply)
	err := c.cc.Invoke(ctx, Sysuser_ListSysuser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCaptchaReply)
	err := c.cc.Invoke(ctx, Sysuser_GetCaptcha_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Sysuser_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, Sysuser_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, Sysuser_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*ChangeStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeStatusReply)
	err := c.cc.Invoke(ctx, Sysuser_ChangeStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePasswordReply)
	err := c.cc.Invoke(ctx, Sysuser_UpdatePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) GetPostInit(ctx context.Context, in *GetPostInitRequest, opts ...grpc.CallOption) (*GetPostInitReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostInitReply)
	err := c.cc.Invoke(ctx, Sysuser_GetPostInit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) GetUserRolePost(ctx context.Context, in *GetUserRolePostRequest, opts ...grpc.CallOption) (*GetUserRolePostReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserRolePostReply)
	err := c.cc.Invoke(ctx, Sysuser_GetUserRolePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysuserClient) GetUserGoogleSecret(ctx context.Context, in *GetUserGoogleSecretRequest, opts ...grpc.CallOption) (*GetUserGoogleSecretReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserGoogleSecretReply)
	err := c.cc.Invoke(ctx, Sysuser_GetUserGoogleSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysuserServer is the server API for Sysuser service.
// All implementations must embed UnimplementedSysuserServer
// for forward compatibility.
//
// 用户管理
type SysuserServer interface {
	// 创建用户
	CreateSysuser(context.Context, *CreateSysuserRequest) (*CreateSysuserReply, error)
	// 更新用户
	UpdateSysuser(context.Context, *UpdateSysuserRequest) (*UpdateSysuserReply, error)
	// 删除用户
	DeleteSysuser(context.Context, *DeleteSysuserRequest) (*DeleteSysuserReply, error)
	// 获取用户
	GetSysuser(context.Context, *GetSysuserRequest) (*GetSysuserReply, error)
	// 用户列表
	ListSysuser(context.Context, *ListSysuserRequest) (*ListSysuserReply, error)
	// 获取验证码
	GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaReply, error)
	// 登入
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// 登出
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// 获取用户权限
	Auth(context.Context, *AuthRequest) (*AuthReply, error)
	// 更新用户状态
	ChangeStatus(context.Context, *ChangeStatusRequest) (*ChangeStatusReply, error)
	// 更新密码
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordReply, error)
	// 获取岗位
	GetPostInit(context.Context, *GetPostInitRequest) (*GetPostInitReply, error)
	// 获取RoPo
	GetUserRolePost(context.Context, *GetUserRolePostRequest) (*GetUserRolePostReply, error)
	// 生成密钥和二维码
	GetUserGoogleSecret(context.Context, *GetUserGoogleSecretRequest) (*GetUserGoogleSecretReply, error)
	mustEmbedUnimplementedSysuserServer()
}

// UnimplementedSysuserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSysuserServer struct{}

func (UnimplementedSysuserServer) CreateSysuser(context.Context, *CreateSysuserRequest) (*CreateSysuserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysuser not implemented")
}
func (UnimplementedSysuserServer) UpdateSysuser(context.Context, *UpdateSysuserRequest) (*UpdateSysuserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysuser not implemented")
}
func (UnimplementedSysuserServer) DeleteSysuser(context.Context, *DeleteSysuserRequest) (*DeleteSysuserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysuser not implemented")
}
func (UnimplementedSysuserServer) GetSysuser(context.Context, *GetSysuserRequest) (*GetSysuserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysuser not implemented")
}
func (UnimplementedSysuserServer) ListSysuser(context.Context, *ListSysuserRequest) (*ListSysuserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSysuser not implemented")
}
func (UnimplementedSysuserServer) GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcha not implemented")
}
func (UnimplementedSysuserServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSysuserServer) Logout(context.Context, *LogoutRequest) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSysuserServer) Auth(context.Context, *AuthRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedSysuserServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*ChangeStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedSysuserServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedSysuserServer) GetPostInit(context.Context, *GetPostInitRequest) (*GetPostInitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostInit not implemented")
}
func (UnimplementedSysuserServer) GetUserRolePost(context.Context, *GetUserRolePostRequest) (*GetUserRolePostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRolePost not implemented")
}
func (UnimplementedSysuserServer) GetUserGoogleSecret(context.Context, *GetUserGoogleSecretRequest) (*GetUserGoogleSecretReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGoogleSecret not implemented")
}
func (UnimplementedSysuserServer) mustEmbedUnimplementedSysuserServer() {}
func (UnimplementedSysuserServer) testEmbeddedByValue()                 {}

// UnsafeSysuserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysuserServer will
// result in compilation errors.
type UnsafeSysuserServer interface {
	mustEmbedUnimplementedSysuserServer()
}

func RegisterSysuserServer(s grpc.ServiceRegistrar, srv SysuserServer) {
	// If the following call pancis, it indicates UnimplementedSysuserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Sysuser_ServiceDesc, srv)
}

func _Sysuser_CreateSysuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSysuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).CreateSysuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_CreateSysuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).CreateSysuser(ctx, req.(*CreateSysuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_UpdateSysuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSysuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).UpdateSysuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_UpdateSysuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).UpdateSysuser(ctx, req.(*UpdateSysuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_DeleteSysuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).DeleteSysuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_DeleteSysuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).DeleteSysuser(ctx, req.(*DeleteSysuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_GetSysuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).GetSysuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_GetSysuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).GetSysuser(ctx, req.(*GetSysuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_ListSysuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).ListSysuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_ListSysuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).ListSysuser(ctx, req.(*ListSysuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_GetCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).GetCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_GetCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).GetCaptcha(ctx, req.(*GetCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_ChangeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_UpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_GetPostInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).GetPostInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_GetPostInit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).GetPostInit(ctx, req.(*GetPostInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_GetUserRolePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRolePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).GetUserRolePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_GetUserRolePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).GetUserRolePost(ctx, req.(*GetUserRolePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sysuser_GetUserGoogleSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGoogleSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysuserServer).GetUserGoogleSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sysuser_GetUserGoogleSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysuserServer).GetUserGoogleSecret(ctx, req.(*GetUserGoogleSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sysuser_ServiceDesc is the grpc.ServiceDesc for Sysuser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sysuser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.http.cas.v1.Sysuser",
	HandlerType: (*SysuserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSysuser",
			Handler:    _Sysuser_CreateSysuser_Handler,
		},
		{
			MethodName: "UpdateSysuser",
			Handler:    _Sysuser_UpdateSysuser_Handler,
		},
		{
			MethodName: "DeleteSysuser",
			Handler:    _Sysuser_DeleteSysuser_Handler,
		},
		{
			MethodName: "GetSysuser",
			Handler:    _Sysuser_GetSysuser_Handler,
		},
		{
			MethodName: "ListSysuser",
			Handler:    _Sysuser_ListSysuser_Handler,
		},
		{
			MethodName: "GetCaptcha",
			Handler:    _Sysuser_GetCaptcha_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Sysuser_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Sysuser_Logout_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Sysuser_Auth_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _Sysuser_ChangeStatus_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Sysuser_UpdatePassword_Handler,
		},
		{
			MethodName: "GetPostInit",
			Handler:    _Sysuser_GetPostInit_Handler,
		},
		{
			MethodName: "GetUserRolePost",
			Handler:    _Sysuser_GetUserRolePost_Handler,
		},
		{
			MethodName: "GetUserGoogleSecret",
			Handler:    _Sysuser_GetUserGoogleSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "http/cas/v1/sysuser.proto",
}
