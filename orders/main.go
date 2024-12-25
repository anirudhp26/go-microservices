package main

import (
	"context"
	"log"
	"net"

	"github.com/anirudhp26/commons"
	pg "github.com/anirudhp26/commons/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddr           = commons.EnvString("GRPC_ADDR", "localhost:5001")
	paymentServiceAddr = commons.EnvString("PAYMENT_SERVICE_ADDR", "localhost:5002")
)

func main() {
	grpcServer := grpc.NewServer()
	store := NewStore()
	svc := NewService(store)

	psConn, err := grpc.NewClient(paymentServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}
	defer psConn.Close()
	paymentServiceClient := pg.NewPaymentServiceClient(psConn)
	NewGRPCHandler(grpcServer, paymentServiceClient)

	svc.CreateOrder(context.Background())

	log.Println("Started gRPC server on", grpcAddr)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
