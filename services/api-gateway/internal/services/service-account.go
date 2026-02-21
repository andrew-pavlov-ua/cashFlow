package services

import (
	pb "github.com/andrew-pavlov-ua/proto/accounts"
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/handler"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer

	Handler *handler.Handler
}

func NewAccountService(handler *handler.Handler) *AccountService {
	return &AccountService{Handler: handler}
}
