package server

import (
	"context"

	pb "github.com/andrew-pavlov-ua/proto/clients" // accounts grpc proto package
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/handler"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedClientServiceServer

	Handler *handler.Handler
}

// NewGRPCServer initializes and returns a new GRPCServer instance
func NewGRPCServer(handler *handler.Handler) *GRPCServer {
	return &GRPCServer{
		Handler: handler,
	}
}

func (s *GRPCServer) RegisterServices(grpcServer *grpc.Server) {
	pb.RegisterClientServiceServer(grpcServer, s)
}

// Start starts the gRPC server
func (s *GRPCServer) Start() {
	// gRPC server start logic can be added here
}

// Stop gracefully stops the gRPC server
func (s *GRPCServer) Stop() {
	// gRPC server stop logic can be added here
}

// Stop gracefully stops the gRPC server
func (s *GRPCServer) InitClient(_ context.Context, mono_token string) (*pb.ClientResponse, error) {
	// Call the handler to initialize the client
	err := s.Handler.InitClient(mono_token)
	if err != nil {
		return &pb.ClientResponse{
			Success: false,
			Message: "Failed to initialize client",
		}, err
	}

	return &pb.ClientResponse{
		Success: true,
		Message: "Client initialized successfully",
	}, nil
}
