package main

import (
	"log"
	"net/http"

	commons "github.com/anirudhp26/commons"
	_ "github.com/joho/godotenv/autoload"
)

var (
	port = commons.EnvString("GATEWAY_PORT", ":3000") // Port on which the gateway will run
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.InitRoutes(mux)

	log.Default().Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
