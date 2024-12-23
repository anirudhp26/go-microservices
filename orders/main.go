package main

import (
	"context"
	"log"
	"net"

	"github.com/anirudhp26/commons"
	"google.golang.org/grpc"
)

var (
	grpcAddr = commons.EnvString("GRPC_ADDR", "localhost:5001")
)

func main() {
	grpcServer := grpc.NewServer()
	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer)

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
