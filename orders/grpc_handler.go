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
}

// NewGRPCHandler creates a new gRPC handler
func NewGRPCHandler(s *grpc.Server) {
	handler := &grpcHandler{}
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
	return order, nil
	// Implement the logic here
}
