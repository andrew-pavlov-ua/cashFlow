package utils

import (
	"os"
	"time"

	"github.com/andrew-pavlov-ua/pkg/logger"
	"github.com/andrew-pavlov-ua/pkg/models"
	pb "github.com/andrew-pavlov-ua/proto/transactions"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func Getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Warnf("Environment variable %s not set, using default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

func TransactionRequestToProto(req *models.TransactionRequest) *pb.Transaction {
	msg := &pb.Transaction{
		ExternalId:      req.ExternalId,
		Amount:          req.Amount,
		CurrencyCode:    req.CurrencyCode,
		Description:     req.Description,
		TransactionTime: timestamppb.New(time.Unix(req.TransactionTime, 0)),
	}

	return msg
}
