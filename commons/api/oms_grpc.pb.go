// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: api/oms.proto

package oms

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
	OrderService_CreateOrder_FullMethodName = "/api.OrderService/CreateOrder"
	OrderService_GetOrder_FullMethodName    = "/api.OrderService/GetOrder"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, OrderService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility.
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *GetOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}
func (UnimplementedOrderServiceServer) testEmbeddedByValue()                      {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/oms.proto",
}

const (
	PaymentService_ProcessPayment_FullMethodName = "/api.PaymentService/ProcessPayment"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Payments service
type PaymentServiceClient interface {
	ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessPaymentResponse)
	err := c.cc.Invoke(ctx, PaymentService_ProcessPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
//
// Payments service
type PaymentServiceServer interface {
	ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) ProcessPayment(context.Context, *ProcessPaymentRequest) (*ProcessPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessPayment not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_ProcessPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_ProcessPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).ProcessPayment(ctx, req.(*ProcessPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessPayment",
			Handler:    _PaymentService_ProcessPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/oms.proto",
}

const (
	StockService_CheckOutStock_FullMethodName = "/api.StockService/CheckOutStock"
	StockService_CheckInStock_FullMethodName  = "/api.StockService/CheckInStock"
)

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Stock service
type StockServiceClient interface {
	CheckOutStock(ctx context.Context, in *CheckOutStockRequest, opts ...grpc.CallOption) (*MessageStatusResponse, error)
	CheckInStock(ctx context.Context, in *CheckInStockRequest, opts ...grpc.CallOption) (*MessageStatusResponse, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) CheckOutStock(ctx context.Context, in *CheckOutStockRequest, opts ...grpc.CallOption) (*MessageStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageStatusResponse)
	err := c.cc.Invoke(ctx, StockService_CheckOutStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) CheckInStock(ctx context.Context, in *CheckInStockRequest, opts ...grpc.CallOption) (*MessageStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageStatusResponse)
	err := c.cc.Invoke(ctx, StockService_CheckInStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility.
//
// Stock service
type StockServiceServer interface {
	CheckOutStock(context.Context, *CheckOutStockRequest) (*MessageStatusResponse, error)
	CheckInStock(context.Context, *CheckInStockRequest) (*MessageStatusResponse, error)
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStockServiceServer struct{}

func (UnimplementedStockServiceServer) CheckOutStock(context.Context, *CheckOutStockRequest) (*MessageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOutStock not implemented")
}
func (UnimplementedStockServiceServer) CheckInStock(context.Context, *CheckInStockRequest) (*MessageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInStock not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}
func (UnimplementedStockServiceServer) testEmbeddedByValue()                      {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	// If the following call pancis, it indicates UnimplementedStockServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_CheckOutStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOutStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CheckOutStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_CheckOutStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CheckOutStock(ctx, req.(*CheckOutStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_CheckInStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CheckInStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StockService_CheckInStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CheckInStock(ctx, req.(*CheckInStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckOutStock",
			Handler:    _StockService_CheckOutStock_Handler,
		},
		{
			MethodName: "CheckInStock",
			Handler:    _StockService_CheckInStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/oms.proto",
}
