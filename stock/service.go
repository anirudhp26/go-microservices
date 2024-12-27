package main

import (
	"context"

	pb "github.com/anirudhp26/commons/api"
	"google.golang.org/grpc"
)

type handler struct {
	// Add any dependencies here
	pb.UnimplementedStockServiceServer
}

func NewGRPCHandler(server *grpc.Server) {
	handler := &handler{}
	pb.RegisterStockServiceServer(server, handler)
}

func (h *handler) CheckOutStock(ctx context.Context, payload *pb.CheckOutStockRequest) (*pb.MessageStatusResponse, error) {
	// Add your implementation here
	return nil, nil
}
