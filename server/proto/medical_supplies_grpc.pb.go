// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: medical_supplies.proto

package medical_supplies

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

// MedicalSuppliesClient is the client API for MedicalSupplies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedicalSuppliesClient interface {
	TestRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DummyResponse, error)
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	FetchUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error)
	UpdateUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	InsertEquipment(ctx context.Context, in *Item, opts ...grpc.CallOption) (*EquipmentResponse, error)
	FetchEquipment(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EquipmentResponse, error)
	UpdateEquipment(ctx context.Context, in *User, opts ...grpc.CallOption) (*EquipmentResponse, error)
	DeleteEquipment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*EquipmentResponse, error)
}

type medicalSuppliesClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicalSuppliesClient(cc grpc.ClientConnInterface) MedicalSuppliesClient {
	return &medicalSuppliesClient{cc}
}

func (c *medicalSuppliesClient) TestRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DummyResponse, error) {
	out := new(DummyResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/TestRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) FetchUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/FetchUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) UpdateUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) InsertEquipment(ctx context.Context, in *Item, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/InsertEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) FetchEquipment(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/FetchEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) UpdateEquipment(ctx context.Context, in *User, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/UpdateEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalSuppliesClient) DeleteEquipment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*EquipmentResponse, error) {
	out := new(EquipmentResponse)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/DeleteEquipment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicalSuppliesServer is the server API for MedicalSupplies service.
// All implementations must embed UnimplementedMedicalSuppliesServer
// for forward compatibility
type MedicalSuppliesServer interface {
	TestRequest(context.Context, *Request) (*DummyResponse, error)
	CreateUser(context.Context, *User) (*UserResponse, error)
	FetchUsers(context.Context, *emptypb.Empty) (*UserResponse, error)
	UpdateUser(context.Context, *Request) (*UserResponse, error)
	DeleteUser(context.Context, *User) (*UserResponse, error)
	InsertEquipment(context.Context, *Item) (*EquipmentResponse, error)
	FetchEquipment(context.Context, *emptypb.Empty) (*EquipmentResponse, error)
	UpdateEquipment(context.Context, *User) (*EquipmentResponse, error)
	DeleteEquipment(context.Context, *Request) (*EquipmentResponse, error)
	mustEmbedUnimplementedMedicalSuppliesServer()
}

// UnimplementedMedicalSuppliesServer must be embedded to have forward compatible implementations.
type UnimplementedMedicalSuppliesServer struct {
}

func (UnimplementedMedicalSuppliesServer) TestRequest(context.Context, *Request) (*DummyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRequest not implemented")
}
func (UnimplementedMedicalSuppliesServer) CreateUser(context.Context, *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMedicalSuppliesServer) FetchUsers(context.Context, *emptypb.Empty) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUsers not implemented")
}
func (UnimplementedMedicalSuppliesServer) UpdateUser(context.Context, *Request) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMedicalSuppliesServer) DeleteUser(context.Context, *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedMedicalSuppliesServer) InsertEquipment(context.Context, *Item) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEquipment not implemented")
}
func (UnimplementedMedicalSuppliesServer) FetchEquipment(context.Context, *emptypb.Empty) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchEquipment not implemented")
}
func (UnimplementedMedicalSuppliesServer) UpdateEquipment(context.Context, *User) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEquipment not implemented")
}
func (UnimplementedMedicalSuppliesServer) DeleteEquipment(context.Context, *Request) (*EquipmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEquipment not implemented")
}
func (UnimplementedMedicalSuppliesServer) mustEmbedUnimplementedMedicalSuppliesServer() {}

// UnsafeMedicalSuppliesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedicalSuppliesServer will
// result in compilation errors.
type UnsafeMedicalSuppliesServer interface {
	mustEmbedUnimplementedMedicalSuppliesServer()
}

func RegisterMedicalSuppliesServer(s grpc.ServiceRegistrar, srv MedicalSuppliesServer) {
	s.RegisterService(&MedicalSupplies_ServiceDesc, srv)
}

func _MedicalSupplies_TestRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).TestRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/TestRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).TestRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_FetchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).FetchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/FetchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).FetchUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).UpdateUser(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_InsertEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).InsertEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/InsertEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).InsertEquipment(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_FetchEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).FetchEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/FetchEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).FetchEquipment(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_UpdateEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).UpdateEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/UpdateEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).UpdateEquipment(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalSupplies_DeleteEquipment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).DeleteEquipment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/DeleteEquipment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).DeleteEquipment(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MedicalSupplies_ServiceDesc is the grpc.ServiceDesc for MedicalSupplies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedicalSupplies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medical_supplies.MedicalSupplies",
	HandlerType: (*MedicalSuppliesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestRequest",
			Handler:    _MedicalSupplies_TestRequest_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _MedicalSupplies_CreateUser_Handler,
		},
		{
			MethodName: "FetchUsers",
			Handler:    _MedicalSupplies_FetchUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _MedicalSupplies_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _MedicalSupplies_DeleteUser_Handler,
		},
		{
			MethodName: "InsertEquipment",
			Handler:    _MedicalSupplies_InsertEquipment_Handler,
		},
		{
			MethodName: "FetchEquipment",
			Handler:    _MedicalSupplies_FetchEquipment_Handler,
		},
		{
			MethodName: "UpdateEquipment",
			Handler:    _MedicalSupplies_UpdateEquipment_Handler,
		},
		{
			MethodName: "DeleteEquipment",
			Handler:    _MedicalSupplies_DeleteEquipment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medical_supplies.proto",
}
