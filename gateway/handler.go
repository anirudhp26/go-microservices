package main

import (
	"fmt"
	"net/http"
)

// Handler is a struct that will hold the routes for the server
type handler struct {
	// Add any dependencies here
}

// Creates a new handler
func NewHandler() *handler {
	return &handler{}
}

// Initializes the routes for the handler
func (h *handler) InitRoutes(mux *http.ServeMux) {
	fmt.Println("Initializing routes")
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/orders", h.CreateOrder)
}

// CreateOrder is a handler for creating an order
func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create order"))
}
