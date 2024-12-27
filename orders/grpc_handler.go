package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/anirudhp26/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	// Add any dependencies here
	pb.UnimplementedOrderServiceServer
	paymentServiceClient pb.PaymentServiceClient
	stockServiceClient   pb.StockServiceClient
}

// NewGRPCHandler creates a new gRPC handler
func NewGRPCHandler(s *grpc.Server, paymentServiceClient pb.PaymentServiceClient, stockServiceClient pb.StockServiceClient) {
	handler := &grpcHandler{paymentServiceClient: paymentServiceClient, stockServiceClient: stockServiceClient}
	pb.RegisterOrderServiceServer(s, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, payload *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order Received for customerID %v and items %v", payload.CustomerId, payload.Items)
	for _, item := range payload.Items {
		log.Printf("Item: %v, Quantity: %v", item.ID, item.Quantity)
	}

	transactionId := fmt.Sprintf("%s-%d", payload.CustomerId, time.Now().Unix())
	order := &pb.Order{
		ID:            "123",
		CustomerId:    payload.CustomerId,
		Status:        "Pending",
		TransactionId: transactionId,
	}
	return order, nil
}
