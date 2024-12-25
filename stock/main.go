package main

import (
	"log"
	"net"

	"github.com/anirudhp26/commons"
	"google.golang.org/grpc"
)

var (
	stockServiceAddr = commons.EnvString("STOCK_SERVICE_ADDR", "localhost:5003")
)

func main() {
	grpcServer := grpc.NewServer()
	defer grpcServer.Stop()

	NewGRPCHandler(grpcServer)

	listner, err := net.Listen("tcp", stockServiceAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
