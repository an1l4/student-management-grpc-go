// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: studentmgmt/studentmgmt.proto

package go_studentmgmt_grpc

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

// StudentManagementClient is the client API for StudentManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentManagementClient interface {
	CreateNewStudent(ctx context.Context, in *StudentList, opts ...grpc.CallOption) (*Id, error)
	GetAllStudent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (StudentManagement_GetAllStudentClient, error)
	GetStudentBySection(ctx context.Context, in *Section, opts ...grpc.CallOption) (*StudentList, error)
	GetStudentByClass(ctx context.Context, in *Class, opts ...grpc.CallOption) (*StudentList, error)
	GetStudentByClassAndSection(ctx context.Context, in *ClassSectionInfo, opts ...grpc.CallOption) (StudentManagement_GetStudentByClassAndSectionClient, error)
	UpdateStudent(ctx context.Context, in *StudentList, opts ...grpc.CallOption) (*Status, error)
	DeleteStudent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error)
}

type studentManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentManagementClient(cc grpc.ClientConnInterface) StudentManagementClient {
	return &studentManagementClient{cc}
}

func (c *studentManagementClient) CreateNewStudent(ctx context.Context, in *StudentList, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/studentmgmt.StudentManagement/CreateNewStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentManagementClient) GetAllStudent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (StudentManagement_GetAllStudentClient, error) {
	stream, err := c.cc.NewStream(ctx, &StudentManagement_ServiceDesc.Streams[0], "/studentmgmt.StudentManagement/GetAllStudent", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentManagementGetAllStudentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StudentManagement_GetAllStudentClient interface {
	Recv() (*StudentList, error)
	grpc.ClientStream
}

type studentManagementGetAllStudentClient struct {
	grpc.ClientStream
}

func (x *studentManagementGetAllStudentClient) Recv() (*StudentList, error) {
	m := new(StudentList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *studentManagementClient) GetStudentBySection(ctx context.Context, in *Section, opts ...grpc.CallOption) (*StudentList, error) {
	out := new(StudentList)
	err := c.cc.Invoke(ctx, "/studentmgmt.StudentManagement/GetStudentBySection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentManagementClient) GetStudentByClass(ctx context.Context, in *Class, opts ...grpc.CallOption) (*StudentList, error) {
	out := new(StudentList)
	err := c.cc.Invoke(ctx, "/studentmgmt.StudentManagement/GetStudentByClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentManagementClient) GetStudentByClassAndSection(ctx context.Context, in *ClassSectionInfo, opts ...grpc.CallOption) (StudentManagement_GetStudentByClassAndSectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &StudentManagement_ServiceDesc.Streams[1], "/studentmgmt.StudentManagement/GetStudentByClassAndSection", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentManagementGetStudentByClassAndSectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StudentManagement_GetStudentByClassAndSectionClient interface {
	Recv() (*StudentList, error)
	grpc.ClientStream
}

type studentManagementGetStudentByClassAndSectionClient struct {
	grpc.ClientStream
}

func (x *studentManagementGetStudentByClassAndSectionClient) Recv() (*StudentList, error) {
	m := new(StudentList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *studentManagementClient) UpdateStudent(ctx context.Context, in *StudentList, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/studentmgmt.StudentManagement/UpdateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentManagementClient) DeleteStudent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/studentmgmt.StudentManagement/DeleteStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentManagementServer is the server API for StudentManagement service.
// All implementations must embed UnimplementedStudentManagementServer
// for forward compatibility
type StudentManagementServer interface {
	CreateNewStudent(context.Context, *StudentList) (*Id, error)
	GetAllStudent(*Empty, StudentManagement_GetAllStudentServer) error
	GetStudentBySection(context.Context, *Section) (*StudentList, error)
	GetStudentByClass(context.Context, *Class) (*StudentList, error)
	GetStudentByClassAndSection(*ClassSectionInfo, StudentManagement_GetStudentByClassAndSectionServer) error
	UpdateStudent(context.Context, *StudentList) (*Status, error)
	DeleteStudent(context.Context, *Id) (*Status, error)
	mustEmbedUnimplementedStudentManagementServer()
}

// UnimplementedStudentManagementServer must be embedded to have forward compatible implementations.
type UnimplementedStudentManagementServer struct {
}

func (UnimplementedStudentManagementServer) CreateNewStudent(context.Context, *StudentList) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewStudent not implemented")
}
func (UnimplementedStudentManagementServer) GetAllStudent(*Empty, StudentManagement_GetAllStudentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllStudent not implemented")
}
func (UnimplementedStudentManagementServer) GetStudentBySection(context.Context, *Section) (*StudentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentBySection not implemented")
}
func (UnimplementedStudentManagementServer) GetStudentByClass(context.Context, *Class) (*StudentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentByClass not implemented")
}
func (UnimplementedStudentManagementServer) GetStudentByClassAndSection(*ClassSectionInfo, StudentManagement_GetStudentByClassAndSectionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStudentByClassAndSection not implemented")
}
func (UnimplementedStudentManagementServer) UpdateStudent(context.Context, *StudentList) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedStudentManagementServer) DeleteStudent(context.Context, *Id) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedStudentManagementServer) mustEmbedUnimplementedStudentManagementServer() {}

// UnsafeStudentManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentManagementServer will
// result in compilation errors.
type UnsafeStudentManagementServer interface {
	mustEmbedUnimplementedStudentManagementServer()
}

func RegisterStudentManagementServer(s grpc.ServiceRegistrar, srv StudentManagementServer) {
	s.RegisterService(&StudentManagement_ServiceDesc, srv)
}

func _StudentManagement_CreateNewStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentManagementServer).CreateNewStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/studentmgmt.StudentManagement/CreateNewStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentManagementServer).CreateNewStudent(ctx, req.(*StudentList))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentManagement_GetAllStudent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentManagementServer).GetAllStudent(m, &studentManagementGetAllStudentServer{stream})
}

type StudentManagement_GetAllStudentServer interface {
	Send(*StudentList) error
	grpc.ServerStream
}

type studentManagementGetAllStudentServer struct {
	grpc.ServerStream
}

func (x *studentManagementGetAllStudentServer) Send(m *StudentList) error {
	return x.ServerStream.SendMsg(m)
}

func _StudentManagement_GetStudentBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Section)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentManagementServer).GetStudentBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/studentmgmt.StudentManagement/GetStudentBySection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentManagementServer).GetStudentBySection(ctx, req.(*Section))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentManagement_GetStudentByClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Class)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentManagementServer).GetStudentByClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/studentmgmt.StudentManagement/GetStudentByClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentManagementServer).GetStudentByClass(ctx, req.(*Class))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentManagement_GetStudentByClassAndSection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClassSectionInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentManagementServer).GetStudentByClassAndSection(m, &studentManagementGetStudentByClassAndSectionServer{stream})
}

type StudentManagement_GetStudentByClassAndSectionServer interface {
	Send(*StudentList) error
	grpc.ServerStream
}

type studentManagementGetStudentByClassAndSectionServer struct {
	grpc.ServerStream
}

func (x *studentManagementGetStudentByClassAndSectionServer) Send(m *StudentList) error {
	return x.ServerStream.SendMsg(m)
}

func _StudentManagement_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentManagementServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/studentmgmt.StudentManagement/UpdateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentManagementServer).UpdateStudent(ctx, req.(*StudentList))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentManagement_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentManagementServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/studentmgmt.StudentManagement/DeleteStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentManagementServer).DeleteStudent(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentManagement_ServiceDesc is the grpc.ServiceDesc for StudentManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "studentmgmt.StudentManagement",
	HandlerType: (*StudentManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewStudent",
			Handler:    _StudentManagement_CreateNewStudent_Handler,
		},
		{
			MethodName: "GetStudentBySection",
			Handler:    _StudentManagement_GetStudentBySection_Handler,
		},
		{
			MethodName: "GetStudentByClass",
			Handler:    _StudentManagement_GetStudentByClass_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _StudentManagement_UpdateStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _StudentManagement_DeleteStudent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllStudent",
			Handler:       _StudentManagement_GetAllStudent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStudentByClassAndSection",
			Handler:       _StudentManagement_GetStudentByClassAndSection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "studentmgmt/studentmgmt.proto",
}
