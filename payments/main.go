package main

import (
	"log"
	"net"

	"github.com/anirudhp26/commons"
	"google.golang.org/grpc"
)

var (
	paymentServiceArrr = commons.EnvString("PAYMENT_SERVICE_ADDR", "localhost:5002")
)

func main() {
	grpcPaymentServer := grpc.NewServer()
	defer grpcPaymentServer.Stop()
	NewGRPCService(grpcPaymentServer) // Initializes the gRPC service

	listner, err := net.Listen("tcp", paymentServiceArrr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listner.Close()
	log.Println("Started gRPC server on", paymentServiceArrr)

	if err := grpcPaymentServer.Serve(listner); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}
}
