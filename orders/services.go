package main

import "context"

type service struct {
	// Add any dependencies here
	store OrderStore
}

// NewService creates a new service
func NewService(s OrderStore) *service {
	return &service{s}
}

func (s *service) CreateOrder(ctx context.Context) error {
	return nil
	// Implement the logic here
}
