package main

import (
	"context"
	"log"

	pb "github.com/anirudhp26/commons/api"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedPaymentServiceServer
}

func NewGRPCService(s *grpc.Server) {
	handler := &service{}
	pb.RegisterPaymentServiceServer(s, handler)
}

func (s *service) ProcessPayment(ctx context.Context, payload *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	// Accept payment logic
	log.Printf("Payment Process requested: %v", payload)
	return &pb.ProcessPaymentResponse{
		Success:   true,
		PaymentId: "123",
	}, nil
}
