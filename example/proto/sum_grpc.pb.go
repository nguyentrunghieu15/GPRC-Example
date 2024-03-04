// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: example/proto_file/sum.proto

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

// CaculateSumClient is the client API for CaculateSum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaculateSumClient interface {
	SumTowInteger(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	SumArr(ctx context.Context, in *SumArrRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (CaculateSum_PrimesClient, error)
	Avg(ctx context.Context, opts ...grpc.CallOption) (CaculateSum_AvgClient, error)
	Max(ctx context.Context, opts ...grpc.CallOption) (CaculateSum_MaxClient, error)
}

type caculateSumClient struct {
	cc grpc.ClientConnInterface
}

func NewCaculateSumClient(cc grpc.ClientConnInterface) CaculateSumClient {
	return &caculateSumClient{cc}
}

func (c *caculateSumClient) SumTowInteger(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/CaculateSum/SumTowInteger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caculateSumClient) SumArr(ctx context.Context, in *SumArrRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/CaculateSum/SumArr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caculateSumClient) Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (CaculateSum_PrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CaculateSum_ServiceDesc.Streams[0], "/CaculateSum/Primes", opts...)
	if err != nil {
		return nil, err
	}
	x := &caculateSumPrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CaculateSum_PrimesClient interface {
	Recv() (*PrimesResponse, error)
	grpc.ClientStream
}

type caculateSumPrimesClient struct {
	grpc.ClientStream
}

func (x *caculateSumPrimesClient) Recv() (*PrimesResponse, error) {
	m := new(PrimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *caculateSumClient) Avg(ctx context.Context, opts ...grpc.CallOption) (CaculateSum_AvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CaculateSum_ServiceDesc.Streams[1], "/CaculateSum/Avg", opts...)
	if err != nil {
		return nil, err
	}
	x := &caculateSumAvgClient{stream}
	return x, nil
}

type CaculateSum_AvgClient interface {
	Send(*PrimesRequest) error
	CloseAndRecv() (*AvgResponse, error)
	grpc.ClientStream
}

type caculateSumAvgClient struct {
	grpc.ClientStream
}

func (x *caculateSumAvgClient) Send(m *PrimesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *caculateSumAvgClient) CloseAndRecv() (*AvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *caculateSumClient) Max(ctx context.Context, opts ...grpc.CallOption) (CaculateSum_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CaculateSum_ServiceDesc.Streams[2], "/CaculateSum/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &caculateSumMaxClient{stream}
	return x, nil
}

type CaculateSum_MaxClient interface {
	Send(*PrimesRequest) error
	Recv() (*PrimesResponse, error)
	grpc.ClientStream
}

type caculateSumMaxClient struct {
	grpc.ClientStream
}

func (x *caculateSumMaxClient) Send(m *PrimesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *caculateSumMaxClient) Recv() (*PrimesResponse, error) {
	m := new(PrimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CaculateSumServer is the server API for CaculateSum service.
// All implementations must embed UnimplementedCaculateSumServer
// for forward compatibility
type CaculateSumServer interface {
	SumTowInteger(context.Context, *SumRequest) (*SumResponse, error)
	SumArr(context.Context, *SumArrRequest) (*SumResponse, error)
	Primes(*PrimesRequest, CaculateSum_PrimesServer) error
	Avg(CaculateSum_AvgServer) error
	Max(CaculateSum_MaxServer) error
	mustEmbedUnimplementedCaculateSumServer()
}

// UnimplementedCaculateSumServer must be embedded to have forward compatible implementations.
type UnimplementedCaculateSumServer struct {
}

func (UnimplementedCaculateSumServer) SumTowInteger(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumTowInteger not implemented")
}
func (UnimplementedCaculateSumServer) SumArr(context.Context, *SumArrRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumArr not implemented")
}
func (UnimplementedCaculateSumServer) Primes(*PrimesRequest, CaculateSum_PrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedCaculateSumServer) Avg(CaculateSum_AvgServer) error {
	return status.Errorf(codes.Unimplemented, "method Avg not implemented")
}
func (UnimplementedCaculateSumServer) Max(CaculateSum_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedCaculateSumServer) mustEmbedUnimplementedCaculateSumServer() {}

// UnsafeCaculateSumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaculateSumServer will
// result in compilation errors.
type UnsafeCaculateSumServer interface {
	mustEmbedUnimplementedCaculateSumServer()
}

func RegisterCaculateSumServer(s grpc.ServiceRegistrar, srv CaculateSumServer) {
	s.RegisterService(&CaculateSum_ServiceDesc, srv)
}

func _CaculateSum_SumTowInteger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaculateSumServer).SumTowInteger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CaculateSum/SumTowInteger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaculateSumServer).SumTowInteger(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaculateSum_SumArr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumArrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaculateSumServer).SumArr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CaculateSum/SumArr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaculateSumServer).SumArr(ctx, req.(*SumArrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaculateSum_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CaculateSumServer).Primes(m, &caculateSumPrimesServer{stream})
}

type CaculateSum_PrimesServer interface {
	Send(*PrimesResponse) error
	grpc.ServerStream
}

type caculateSumPrimesServer struct {
	grpc.ServerStream
}

func (x *caculateSumPrimesServer) Send(m *PrimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CaculateSum_Avg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CaculateSumServer).Avg(&caculateSumAvgServer{stream})
}

type CaculateSum_AvgServer interface {
	SendAndClose(*AvgResponse) error
	Recv() (*PrimesRequest, error)
	grpc.ServerStream
}

type caculateSumAvgServer struct {
	grpc.ServerStream
}

func (x *caculateSumAvgServer) SendAndClose(m *AvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *caculateSumAvgServer) Recv() (*PrimesRequest, error) {
	m := new(PrimesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CaculateSum_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CaculateSumServer).Max(&caculateSumMaxServer{stream})
}

type CaculateSum_MaxServer interface {
	Send(*PrimesResponse) error
	Recv() (*PrimesRequest, error)
	grpc.ServerStream
}

type caculateSumMaxServer struct {
	grpc.ServerStream
}

func (x *caculateSumMaxServer) Send(m *PrimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *caculateSumMaxServer) Recv() (*PrimesRequest, error) {
	m := new(PrimesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CaculateSum_ServiceDesc is the grpc.ServiceDesc for CaculateSum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaculateSum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CaculateSum",
	HandlerType: (*CaculateSumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SumTowInteger",
			Handler:    _CaculateSum_SumTowInteger_Handler,
		},
		{
			MethodName: "SumArr",
			Handler:    _CaculateSum_SumArr_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _CaculateSum_Primes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Avg",
			Handler:       _CaculateSum_Avg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _CaculateSum_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example/proto_file/sum.proto",
}
