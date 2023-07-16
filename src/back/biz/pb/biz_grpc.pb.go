// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pb/biz.proto

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

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	SendUser(ctx context.Context, in *UserResponse, opts ...grpc.CallOption) (*UserRequest, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	SendGroup(ctx context.Context, in *GroupResponse, opts ...grpc.CallOption) (*GroupRequest, error)
	DeleteGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetSubgroup(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error)
	SendSubgroup(ctx context.Context, in *SubResponse, opts ...grpc.CallOption) (*SubRequest, error)
	DeleteSubgroup(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseResponse, error)
	SendCourse(ctx context.Context, in *CourseResponse, opts ...grpc.CallOption) (*CourseRequest, error)
	DeleteCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	GetAllGroups(ctx context.Context, in *GetAllGroupsRequest, opts ...grpc.CallOption) (*GetAllGroupsResponse, error)
	GetAllCourses(ctx context.Context, in *GetAllCoursesRequest, opts ...grpc.CallOption) (*GetAllCoursesResponse, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) SendUser(ctx context.Context, in *UserResponse, opts ...grpc.CallOption) (*UserRequest, error) {
	out := new(UserRequest)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/SendUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) SendGroup(ctx context.Context, in *GroupResponse, opts ...grpc.CallOption) (*GroupRequest, error) {
	out := new(GroupRequest)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/SendGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) DeleteGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetSubgroup(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubResponse, error) {
	out := new(SubResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/GetSubgroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) SendSubgroup(ctx context.Context, in *SubResponse, opts ...grpc.CallOption) (*SubRequest, error) {
	out := new(SubRequest)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/SendSubgroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) DeleteSubgroup(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/DeleteSubgroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseResponse, error) {
	out := new(CourseResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) SendCourse(ctx context.Context, in *CourseResponse, opts ...grpc.CallOption) (*CourseRequest, error) {
	out := new(CourseRequest)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/SendCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) DeleteCourse(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetAllGroups(ctx context.Context, in *GetAllGroupsRequest, opts ...grpc.CallOption) (*GetAllGroupsResponse, error) {
	out := new(GetAllGroupsResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/GetAllGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetAllCourses(ctx context.Context, in *GetAllCoursesRequest, opts ...grpc.CallOption) (*GetAllCoursesResponse, error) {
	out := new(GetAllCoursesResponse)
	err := c.cc.Invoke(ctx, "/biz.v1.DatabaseService/GetAllCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations must embed UnimplementedDatabaseServiceServer
// for forward compatibility
type DatabaseServiceServer interface {
	GetUser(context.Context, *UserRequest) (*UserResponse, error)
	SendUser(context.Context, *UserResponse) (*UserRequest, error)
	DeleteUser(context.Context, *UserRequest) (*DeleteResponse, error)
	GetGroup(context.Context, *GroupRequest) (*GroupResponse, error)
	SendGroup(context.Context, *GroupResponse) (*GroupRequest, error)
	DeleteGroup(context.Context, *GroupRequest) (*DeleteResponse, error)
	GetSubgroup(context.Context, *SubRequest) (*SubResponse, error)
	SendSubgroup(context.Context, *SubResponse) (*SubRequest, error)
	DeleteSubgroup(context.Context, *SubRequest) (*DeleteResponse, error)
	GetCourse(context.Context, *CourseRequest) (*CourseResponse, error)
	SendCourse(context.Context, *CourseResponse) (*CourseRequest, error)
	DeleteCourse(context.Context, *CourseRequest) (*DeleteResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	GetAllGroups(context.Context, *GetAllGroupsRequest) (*GetAllGroupsResponse, error)
	GetAllCourses(context.Context, *GetAllCoursesRequest) (*GetAllCoursesResponse, error)
	mustEmbedUnimplementedDatabaseServiceServer()
}

// UnimplementedDatabaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (UnimplementedDatabaseServiceServer) GetUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedDatabaseServiceServer) SendUser(context.Context, *UserResponse) (*UserRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUser not implemented")
}
func (UnimplementedDatabaseServiceServer) DeleteUser(context.Context, *UserRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedDatabaseServiceServer) GetGroup(context.Context, *GroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedDatabaseServiceServer) SendGroup(context.Context, *GroupResponse) (*GroupRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroup not implemented")
}
func (UnimplementedDatabaseServiceServer) DeleteGroup(context.Context, *GroupRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedDatabaseServiceServer) GetSubgroup(context.Context, *SubRequest) (*SubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubgroup not implemented")
}
func (UnimplementedDatabaseServiceServer) SendSubgroup(context.Context, *SubResponse) (*SubRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSubgroup not implemented")
}
func (UnimplementedDatabaseServiceServer) DeleteSubgroup(context.Context, *SubRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubgroup not implemented")
}
func (UnimplementedDatabaseServiceServer) GetCourse(context.Context, *CourseRequest) (*CourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedDatabaseServiceServer) SendCourse(context.Context, *CourseResponse) (*CourseRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCourse not implemented")
}
func (UnimplementedDatabaseServiceServer) DeleteCourse(context.Context, *CourseRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedDatabaseServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedDatabaseServiceServer) GetAllGroups(context.Context, *GetAllGroupsRequest) (*GetAllGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGroups not implemented")
}
func (UnimplementedDatabaseServiceServer) GetAllCourses(context.Context, *GetAllCoursesRequest) (*GetAllCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCourses not implemented")
}
func (UnimplementedDatabaseServiceServer) mustEmbedUnimplementedDatabaseServiceServer() {}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_SendUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).SendUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/SendUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).SendUser(ctx, req.(*UserResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_SendGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).SendGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/SendGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).SendGroup(ctx, req.(*GroupResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).DeleteGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetSubgroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetSubgroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/GetSubgroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetSubgroup(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_SendSubgroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).SendSubgroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/SendSubgroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).SendSubgroup(ctx, req.(*SubResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_DeleteSubgroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).DeleteSubgroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/DeleteSubgroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).DeleteSubgroup(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetCourse(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_SendCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).SendCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/SendCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).SendCourse(ctx, req.(*CourseResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).DeleteCourse(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetAllGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetAllGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/GetAllGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetAllGroups(ctx, req.(*GetAllGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetAllCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetAllCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biz.v1.DatabaseService/GetAllCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetAllCourses(ctx, req.(*GetAllCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "biz.v1.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _DatabaseService_GetUser_Handler,
		},
		{
			MethodName: "SendUser",
			Handler:    _DatabaseService_SendUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _DatabaseService_DeleteUser_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _DatabaseService_GetGroup_Handler,
		},
		{
			MethodName: "SendGroup",
			Handler:    _DatabaseService_SendGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _DatabaseService_DeleteGroup_Handler,
		},
		{
			MethodName: "GetSubgroup",
			Handler:    _DatabaseService_GetSubgroup_Handler,
		},
		{
			MethodName: "SendSubgroup",
			Handler:    _DatabaseService_SendSubgroup_Handler,
		},
		{
			MethodName: "DeleteSubgroup",
			Handler:    _DatabaseService_DeleteSubgroup_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _DatabaseService_GetCourse_Handler,
		},
		{
			MethodName: "SendCourse",
			Handler:    _DatabaseService_SendCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _DatabaseService_DeleteCourse_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _DatabaseService_GetAllUsers_Handler,
		},
		{
			MethodName: "GetAllGroups",
			Handler:    _DatabaseService_GetAllGroups_Handler,
		},
		{
			MethodName: "GetAllCourses",
			Handler:    _DatabaseService_GetAllCourses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/biz.proto",
}
