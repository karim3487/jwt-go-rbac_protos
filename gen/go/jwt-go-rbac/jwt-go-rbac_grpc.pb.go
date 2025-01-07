// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: jwt-go-rbac/jwt-go-rbac.proto

package jwt_go_rbacv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_CreateUser_FullMethodName               = "/auth.AuthService/CreateUser"
	AuthService_RegisterUser_FullMethodName             = "/auth.AuthService/RegisterUser"
	AuthService_Login_FullMethodName                    = "/auth.AuthService/Login"
	AuthService_UpdateUser_FullMethodName               = "/auth.AuthService/UpdateUser"
	AuthService_GetUserInfo_FullMethodName              = "/auth.AuthService/GetUserInfo"
	AuthService_GetAllUsers_FullMethodName              = "/auth.AuthService/GetAllUsers"
	AuthService_DeleteUser_FullMethodName               = "/auth.AuthService/DeleteUser"
	AuthService_CreateRole_FullMethodName               = "/auth.AuthService/CreateRole"
	AuthService_UpdateRole_FullMethodName               = "/auth.AuthService/UpdateRole"
	AuthService_DeleteRole_FullMethodName               = "/auth.AuthService/DeleteRole"
	AuthService_GetAllRoles_FullMethodName              = "/auth.AuthService/GetAllRoles"
	AuthService_AddRoleToUser_FullMethodName            = "/auth.AuthService/AddRoleToUser"
	AuthService_RemoveRoleFromUser_FullMethodName       = "/auth.AuthService/RemoveRoleFromUser"
	AuthService_GetRolesForUser_FullMethodName          = "/auth.AuthService/GetRolesForUser"
	AuthService_CreatePermission_FullMethodName         = "/auth.AuthService/CreatePermission"
	AuthService_UpdatePermission_FullMethodName         = "/auth.AuthService/UpdatePermission"
	AuthService_DeletePermission_FullMethodName         = "/auth.AuthService/DeletePermission"
	AuthService_GetAllPermissions_FullMethodName        = "/auth.AuthService/GetAllPermissions"
	AuthService_AddPermissionToRole_FullMethodName      = "/auth.AuthService/AddPermissionToRole"
	AuthService_RemovePermissionFromRole_FullMethodName = "/auth.AuthService/RemovePermissionFromRole"
	AuthService_GetPermissionsForRole_FullMethodName    = "/auth.AuthService/GetPermissionsForRole"
	AuthService_CheckPermissions_FullMethodName         = "/auth.AuthService/CheckPermissions"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// -------------------- USER MANAGEMENT -------------------------
	CreateUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error)
	GetAllUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserInfoList, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// -------------------- ROLE MANAGEMENT -------------------------
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateOrUpdateRoleResponse, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*CreateOrUpdateRoleResponse, error)
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllRoles(ctx context.Context, in *GetAllRolesRequest, opts ...grpc.CallOption) (*RoleList, error)
	AddRoleToUser(ctx context.Context, in *AddRoleToUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRolesForUser(ctx context.Context, in *GetRolesForUserRequest, opts ...grpc.CallOption) (*RoleList, error)
	// -------------------- PERMISSION MANAGEMENT -------------------
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreateOrUpdatePermissionResponse, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*CreateOrUpdatePermissionResponse, error)
	DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllPermissions(ctx context.Context, in *GetAllPermissionsRequest, opts ...grpc.CallOption) (*PermissionList, error)
	AddPermissionToRole(ctx context.Context, in *AddPermissionToRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPermissionsForRole(ctx context.Context, in *GetPermissionsForRoleRequest, opts ...grpc.CallOption) (*PermissionList, error)
	CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, opts ...grpc.CallOption) (*CheckPermissionsResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CreateUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, AuthService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, AuthService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, AuthService_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, AuthService_GetUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAllUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserInfoList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfoList)
	err := c.cc.Invoke(ctx, AuthService_GetAllUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateOrUpdateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrUpdateRoleResponse)
	err := c.cc.Invoke(ctx, AuthService_CreateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*CreateOrUpdateRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrUpdateRoleResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdateRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_DeleteRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAllRoles(ctx context.Context, in *GetAllRolesRequest, opts ...grpc.CallOption) (*RoleList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleList)
	err := c.cc.Invoke(ctx, AuthService_GetAllRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddRoleToUser(ctx context.Context, in *AddRoleToUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_AddRoleToUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_RemoveRoleFromUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetRolesForUser(ctx context.Context, in *GetRolesForUserRequest, opts ...grpc.CallOption) (*RoleList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleList)
	err := c.cc.Invoke(ctx, AuthService_GetRolesForUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreateOrUpdatePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrUpdatePermissionResponse)
	err := c.cc.Invoke(ctx, AuthService_CreatePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*CreateOrUpdatePermissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrUpdatePermissionResponse)
	err := c.cc.Invoke(ctx, AuthService_UpdatePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeletePermission(ctx context.Context, in *DeletePermissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_DeletePermission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAllPermissions(ctx context.Context, in *GetAllPermissionsRequest, opts ...grpc.CallOption) (*PermissionList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PermissionList)
	err := c.cc.Invoke(ctx, AuthService_GetAllPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddPermissionToRole(ctx context.Context, in *AddPermissionToRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_AddPermissionToRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RemovePermissionFromRole(ctx context.Context, in *RemovePermissionFromRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthService_RemovePermissionFromRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetPermissionsForRole(ctx context.Context, in *GetPermissionsForRoleRequest, opts ...grpc.CallOption) (*PermissionList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PermissionList)
	err := c.cc.Invoke(ctx, AuthService_GetPermissionsForRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, opts ...grpc.CallOption) (*CheckPermissionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckPermissionsResponse)
	err := c.cc.Invoke(ctx, AuthService_CheckPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	// -------------------- USER MANAGEMENT -------------------------
	CreateUser(context.Context, *RegisterUserRequest) (*UserInfo, error)
	RegisterUser(context.Context, *RegisterUserRequest) (*UserInfo, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserInfo, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*UserInfo, error)
	GetAllUsers(context.Context, *emptypb.Empty) (*UserInfoList, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	// -------------------- ROLE MANAGEMENT -------------------------
	CreateRole(context.Context, *CreateRoleRequest) (*CreateOrUpdateRoleResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*CreateOrUpdateRoleResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error)
	GetAllRoles(context.Context, *GetAllRolesRequest) (*RoleList, error)
	AddRoleToUser(context.Context, *AddRoleToUserRequest) (*emptypb.Empty, error)
	RemoveRoleFromUser(context.Context, *RemoveRoleFromUserRequest) (*emptypb.Empty, error)
	GetRolesForUser(context.Context, *GetRolesForUserRequest) (*RoleList, error)
	// -------------------- PERMISSION MANAGEMENT -------------------
	CreatePermission(context.Context, *CreatePermissionRequest) (*CreateOrUpdatePermissionResponse, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*CreateOrUpdatePermissionResponse, error)
	DeletePermission(context.Context, *DeletePermissionRequest) (*emptypb.Empty, error)
	GetAllPermissions(context.Context, *GetAllPermissionsRequest) (*PermissionList, error)
	AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*emptypb.Empty, error)
	RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*emptypb.Empty, error)
	GetPermissionsForRole(context.Context, *GetPermissionsForRoleRequest) (*PermissionList, error)
	CheckPermissions(context.Context, *CheckPermissionsRequest) (*CheckPermissionsResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) CreateUser(context.Context, *RegisterUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAuthServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedAuthServiceServer) GetAllUsers(context.Context, *emptypb.Empty) (*UserInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedAuthServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateOrUpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedAuthServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*CreateOrUpdateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedAuthServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedAuthServiceServer) GetAllRoles(context.Context, *GetAllRolesRequest) (*RoleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoles not implemented")
}
func (UnimplementedAuthServiceServer) AddRoleToUser(context.Context, *AddRoleToUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoleToUser not implemented")
}
func (UnimplementedAuthServiceServer) RemoveRoleFromUser(context.Context, *RemoveRoleFromUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoleFromUser not implemented")
}
func (UnimplementedAuthServiceServer) GetRolesForUser(context.Context, *GetRolesForUserRequest) (*RoleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesForUser not implemented")
}
func (UnimplementedAuthServiceServer) CreatePermission(context.Context, *CreatePermissionRequest) (*CreateOrUpdatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedAuthServiceServer) UpdatePermission(context.Context, *UpdatePermissionRequest) (*CreateOrUpdatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermission not implemented")
}
func (UnimplementedAuthServiceServer) DeletePermission(context.Context, *DeletePermissionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePermission not implemented")
}
func (UnimplementedAuthServiceServer) GetAllPermissions(context.Context, *GetAllPermissionsRequest) (*PermissionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPermissions not implemented")
}
func (UnimplementedAuthServiceServer) AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermissionToRole not implemented")
}
func (UnimplementedAuthServiceServer) RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermissionFromRole not implemented")
}
func (UnimplementedAuthServiceServer) GetPermissionsForRole(context.Context, *GetPermissionsForRoleRequest) (*PermissionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionsForRole not implemented")
}
func (UnimplementedAuthServiceServer) CheckPermissions(context.Context, *CheckPermissionsRequest) (*CheckPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermissions not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAllUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAllRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAllRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetAllRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAllRoles(ctx, req.(*GetAllRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddRoleToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddRoleToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AddRoleToUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddRoleToUser(ctx, req.(*AddRoleToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RemoveRoleFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleFromUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RemoveRoleFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RemoveRoleFromUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RemoveRoleFromUser(ctx, req.(*RemoveRoleFromUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetRolesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetRolesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetRolesForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetRolesForUser(ctx, req.(*GetRolesForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CreatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_UpdatePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeletePermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeletePermission(ctx, req.(*DeletePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAllPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAllPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetAllPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAllPermissions(ctx, req.(*GetAllPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddPermissionToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionToRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddPermissionToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AddPermissionToRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddPermissionToRole(ctx, req.(*AddPermissionToRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RemovePermissionFromRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePermissionFromRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RemovePermissionFromRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RemovePermissionFromRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RemovePermissionFromRole(ctx, req.(*RemovePermissionFromRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetPermissionsForRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsForRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetPermissionsForRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetPermissionsForRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetPermissionsForRole(ctx, req.(*GetPermissionsForRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CheckPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckPermissions(ctx, req.(*CheckPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AuthService_CreateUser_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _AuthService_RegisterUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _AuthService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _AuthService_GetAllUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthService_DeleteUser_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _AuthService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _AuthService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _AuthService_DeleteRole_Handler,
		},
		{
			MethodName: "GetAllRoles",
			Handler:    _AuthService_GetAllRoles_Handler,
		},
		{
			MethodName: "AddRoleToUser",
			Handler:    _AuthService_AddRoleToUser_Handler,
		},
		{
			MethodName: "RemoveRoleFromUser",
			Handler:    _AuthService_RemoveRoleFromUser_Handler,
		},
		{
			MethodName: "GetRolesForUser",
			Handler:    _AuthService_GetRolesForUser_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _AuthService_CreatePermission_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _AuthService_UpdatePermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _AuthService_DeletePermission_Handler,
		},
		{
			MethodName: "GetAllPermissions",
			Handler:    _AuthService_GetAllPermissions_Handler,
		},
		{
			MethodName: "AddPermissionToRole",
			Handler:    _AuthService_AddPermissionToRole_Handler,
		},
		{
			MethodName: "RemovePermissionFromRole",
			Handler:    _AuthService_RemovePermissionFromRole_Handler,
		},
		{
			MethodName: "GetPermissionsForRole",
			Handler:    _AuthService_GetPermissionsForRole_Handler,
		},
		{
			MethodName: "CheckPermissions",
			Handler:    _AuthService_CheckPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwt-go-rbac/jwt-go-rbac.proto",
}
