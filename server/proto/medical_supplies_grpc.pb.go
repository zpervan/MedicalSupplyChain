// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: medical_supplies.proto

package medical_supplies

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

// MedicalSuppliesClient is the client API for MedicalSupplies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedicalSuppliesClient interface {
	FetchAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type medicalSuppliesClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicalSuppliesClient(cc grpc.ClientConnInterface) MedicalSuppliesClient {
	return &medicalSuppliesClient{cc}
}

func (c *medicalSuppliesClient) FetchAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/medical_supplies.MedicalSupplies/FetchAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicalSuppliesServer is the server API for MedicalSupplies service.
// All implementations must embed UnimplementedMedicalSuppliesServer
// for forward compatibility
type MedicalSuppliesServer interface {
	FetchAll(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedMedicalSuppliesServer()
}

// UnimplementedMedicalSuppliesServer must be embedded to have forward compatible implementations.
type UnimplementedMedicalSuppliesServer struct {
}

func (UnimplementedMedicalSuppliesServer) FetchAll(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAll not implemented")
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

func _MedicalSupplies_FetchAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalSuppliesServer).FetchAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical_supplies.MedicalSupplies/FetchAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalSuppliesServer).FetchAll(ctx, req.(*Request))
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
			MethodName: "FetchAll",
			Handler:    _MedicalSupplies_FetchAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medical_supplies.proto",
}
