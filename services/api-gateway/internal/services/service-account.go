package services

import (
	pb "github.com/andrew-pavlov-ua/proto/clients" // accounts grpc proto package
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/handler"
)

type AccountService struct {
	pb.UnimplementedClientServiceServer

	Handler *handler.Handler
}

func NewAccountService(handler *handler.Handler) *AccountService {
	return &AccountService{Handler: handler}
}
