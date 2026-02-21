package server

import (
	pb "github.com/andrew-pavlov-ua/proto/accounts" // accounts grpc proto package
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/handler"
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/services"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	AccountService *services.AccountService
}

// NewGRPCServer initializes and returns a new GRPCServer instance
func NewGRPCServer(handler *handler.Handler) *GRPCServer {
	return &GRPCServer{
		AccountService: services.NewAccountService(handler),
	}
}

func (s *GRPCServer) RegisterServices(grpcServer *grpc.Server) {
	pb.RegisterAccountServiceServer(grpcServer, s.AccountService)
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
func (s *GRPCServer) InitUser() {

}
