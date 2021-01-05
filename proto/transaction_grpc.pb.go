// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProcessTransactionClient is the client API for ProcessTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessTransactionClient interface {
	MakeTransaction(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (ProcessTransaction_MakeTransactionClient, error)
}

type processTransactionClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessTransactionClient(cc grpc.ClientConnInterface) ProcessTransactionClient {
	return &processTransactionClient{cc}
}

func (c *processTransactionClient) MakeTransaction(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (ProcessTransaction_MakeTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessTransaction_serviceDesc.Streams[0], "/proto.ProcessTransaction/MakeTransaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &processTransactionMakeTransactionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessTransaction_MakeTransactionClient interface {
	Recv() (*MoneyResponse, error)
	grpc.ClientStream
}

type processTransactionMakeTransactionClient struct {
	grpc.ClientStream
}

func (x *processTransactionMakeTransactionClient) Recv() (*MoneyResponse, error) {
	m := new(MoneyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessTransactionServer is the server API for ProcessTransaction service.
// All implementations must embed UnimplementedProcessTransactionServer
// for forward compatibility
type ProcessTransactionServer interface {
	MakeTransaction(*MoneyRequest, ProcessTransaction_MakeTransactionServer) error
	mustEmbedUnimplementedProcessTransactionServer()
}

// UnimplementedProcessTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedProcessTransactionServer struct {
}

func (UnimplementedProcessTransactionServer) MakeTransaction(*MoneyRequest, ProcessTransaction_MakeTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method MakeTransaction not implemented")
}
func (UnimplementedProcessTransactionServer) mustEmbedUnimplementedProcessTransactionServer() {}

// UnsafeProcessTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessTransactionServer will
// result in compilation errors.
type UnsafeProcessTransactionServer interface {
	mustEmbedUnimplementedProcessTransactionServer()
}

func RegisterProcessTransactionServer(s grpc.ServiceRegistrar, srv ProcessTransactionServer) {
	s.RegisterService(&_ProcessTransaction_serviceDesc, srv)
}

func _ProcessTransaction_MakeTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MoneyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessTransactionServer).MakeTransaction(m, &processTransactionMakeTransactionServer{stream})
}

type ProcessTransaction_MakeTransactionServer interface {
	Send(*MoneyResponse) error
	grpc.ServerStream
}

type processTransactionMakeTransactionServer struct {
	grpc.ServerStream
}

func (x *processTransactionMakeTransactionServer) Send(m *MoneyResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ProcessTransaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProcessTransaction",
	HandlerType: (*ProcessTransactionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeTransaction",
			Handler:       _ProcessTransaction_MakeTransaction_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transaction.proto",
}