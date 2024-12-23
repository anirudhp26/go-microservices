package main

import (
	"fmt"
	"net/http"

	commons "github.com/anirudhp26/commons"
	pb "github.com/anirudhp26/commons/api"
)

// Handler is a struct that will hold the routes for the server
type handler struct {
	// Add any dependencies here
	orderClient pb.OrderServiceClient
}

// Creates a new handler
func NewHandler(orderClient pb.OrderServiceClient) *handler {
	return &handler{orderClient}
}

// Initializes the routes for the handler
func (h *handler) InitRoutes(mux *http.ServeMux) {
	fmt.Println("Initializing routes")
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/api/v1/customers/{customerId}/orders", h.CreateOrder)
}

// CreateOrder is a handler for creating an order
func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("customerId")             // Get the customer ID from the URL
	var items []*pb.ItemsWithQuantity                   // Create a slice of items
	if err := commons.ReadJson(r, &items); err != nil { // Read the JSON from the request body
		commons.WriteError(w, http.StatusBadRequest, err.Error()) // Write an error response if there was an error
		return
	}
	h.orderClient.CreateOrder(r.Context(), &pb.CreateOrderRequest{ // Call the CreateOrder method on the order service
		CustomerId: customerId,
		Items:      items,
	})
}
