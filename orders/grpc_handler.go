package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/anirudhp26/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	paymentServiceClient pb.PaymentServiceClient
	stockServiceClient   pb.StockServiceClient
}

// NewGRPCHandler creates a new gRPC handler
func NewGRPCHandler(s *grpc.Server, paymentServiceClient pb.PaymentServiceClient, stockServiceClient pb.StockServiceClient) {
	handler := &grpcHandler{paymentServiceClient: paymentServiceClient, stockServiceClient: stockServiceClient}
	pb.RegisterOrderServiceServer(s, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, payload *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order Received for customerID %v and items %v", payload.CustomerId, payload.Items)
	for _, item := range payload.Items {
		log.Printf("Item: %v, Quantity: %v", item.ID, item.Quantity)
	}

	transactionId := fmt.Sprintf("%s-%d", payload.CustomerId, time.Now().Unix())
	order := &pb.Order{
		ID:            "123",
		CustomerId:    payload.CustomerId,
		Status:        "Pending",
		TransactionId: transactionId,
	}
	return order, nil
}

func (h *grpcHandler) ProcessOrder(ctx context.Context, payload *pb.ProcessOrderRequest) (*pb.MessageStatusResponse, error) {
	log.Printf("Processing order: %v", payload)
	paymentRes, paymentErr := h.paymentServiceClient.ProcessPayment(ctx, &pb.ProcessPaymentRequest{
		TransactionId: payload.TransactionId,
		OrderId:       payload.OrderId,
	})
	if paymentErr != nil {
		return nil, paymentErr
	}
	if paymentRes.Success {
		log.Printf("Payment Done for order Id: %v", payload.OrderId)
		stockRes, stockErr := h.stockServiceClient.CheckOutStock(ctx, &pb.CheckOutStockRequest{
			OrderId:    payload.OrderId,
			CustomerId: payload.CustomerId,
			Items:      payload.Items,
		})
		if stockErr != nil {
			return nil, paymentErr
		}
		if stockRes.Success {
			return &pb.MessageStatusResponse{
				Message: "Order Accepted",
				Success: true,
			}, nil
		} else {
			return nil, fmt.Errorf("Some Error occured, try again later")
		}
	} else {
		return nil, fmt.Errorf("Some error Occured, try again later")
	}
}
