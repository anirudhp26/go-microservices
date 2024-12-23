package main

import (
	"log"
	"net/http"

	commons "github.com/anirudhp26/commons"
	pb "github.com/anirudhp26/commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port             = commons.EnvString("GATEWAY_PORT", ":3000")                // Port on which the gateway will run
	orderServiceAddr = commons.EnvString("ORDER_SERVICE_ADDR", "localhost:5001") // Address of the order service
)

func main() {
	// Connection to the order service
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the order service: %v", err)
	}
	log.Printf("Connected to order service at %v", orderServiceAddr)
	defer conn.Close() // Close the connection when the main function exits
	orderclient := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(orderclient)
	handler.InitRoutes(mux)

	log.Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
