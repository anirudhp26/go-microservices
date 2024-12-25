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
	// Add any dependencies here
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
	order := &pb.Order{
		ID: "1",
	}
	// Call the payment service
	transactionId := fmt.Sprintf("%s-%d", payload.CustomerId, time.Now().Unix())
	log.Println("Transaction ID:", transactionId)
	paytmentRes, paymentErr := h.paymentServiceClient.ProcessPayment(ctx, &pb.ProcessPaymentRequest{
		CustomerId:      payload.CustomerId,
		Amount:          100,
		PaymentMethod:   "UPI",
		PaymentMethodId: "2",
		OrderId:         order.ID,
	})
	if paymentErr != nil {
		log.Printf("Error processing payment: %v", paymentErr)
		return nil, paymentErr
	}
	if paytmentRes.Success {
		log.Printf("Payment processed: %v", paytmentRes)
		stockRes, stockErr := h.stockServiceClient.CheckOutStock(ctx, &pb.CheckOutStockRequest{
			Items: payload.Items,
		})
		if stockErr != nil {
			log.Printf("Error checking out stock: %v", stockErr)
			return nil, stockErr
		}
		if stockRes.Success {
			log.Printf("Stock checked out: %v", stockRes)
		}
		return order, nil
	}
	return nil, nil
}
