package grpc

import (
	"context"
	"log"
	"net"

	pb "payment_service/internal/paymentpb"
	"payment_service/internal/service"

	// "google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedPaymentServiceServer
	service service.PaymentServiceServer
}

func NewGRPCServer(service service.PaymentServiceServer) *GRPCServer {
	return &GRPCServer{service: service}
}

func (s *GRPCServer) GeneratePaymentLink(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	return s.service.GeneratePaymentLink(ctx, req)
}

func StartGRPCServer(grpcService service.PaymentServiceServer) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, NewGRPCServer(grpcService))

	// // Включаем reflection API
	// reflection.Register(grpcServer)

	log.Println("gRPC Server is running on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
