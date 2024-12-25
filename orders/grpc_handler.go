package main

import (
	"context"
	"log"

	pb "github.com/anirudhp26/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	// Add any dependencies here
	pb.UnimplementedOrderServiceServer
	paymentServiceClient pb.PaymentServiceClient
}

// NewGRPCHandler creates a new gRPC handler
func NewGRPCHandler(s *grpc.Server, paymentServiceClient pb.PaymentServiceClient) {
	handler := &grpcHandler{paymentServiceClient: paymentServiceClient}
	pb.RegisterOrderServiceServer(s, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, payload *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order Received for customerID %v and items %v", payload.CustomerId, payload.Items)
	for _, item := range payload.Items {
		log.Printf("Item: %v, Quantity: %v", item.ID, item.Quantity)
	}
	order := &pb.Order{
		ID: "1",
	}

	// Call the payment service
	res, err := h.paymentServiceClient.ProcessPayment(ctx, &pb.ProcessPaymentRequest{
		CustomerId:      payload.CustomerId,
		Amount:          100,
		PaymentMethod:   "UPI",
		PaymentMethodId: "2",
		OrderId:         order.ID,
	})
	if err != nil {
		log.Printf("Error processing payment: %v", err)
		return nil, err
	}
	log.Printf("Payment processed: %v", res)

	return order, nil
	// Implement the logic here
}
