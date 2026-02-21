package utils

import (
	"os"
	"time"

	"github.com/andrew-pavlov-ua/pkg/logger"
	"github.com/andrew-pavlov-ua/pkg/models"
	pbc "github.com/andrew-pavlov-ua/proto/clients"
	pbt "github.com/andrew-pavlov-ua/proto/transactions"

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

func TransactionRequestToProto(req *models.TransactionRequest) *pbt.Transaction {
	msg := &pbt.Transaction{
		ExternalId:      req.ExternalId,
		Amount:          req.Amount,
		CurrencyCode:    req.CurrencyCode,
		Description:     req.Description,
		TransactionTime: timestamppb.New(time.Unix(req.TransactionTime, 0)),
	}

	return msg
}

func ClientToProto(client *models.Client) *pbc.Client {
	protoAccs := make([]*pbc.Account, len(client.Accounts))
	for i, acc := range client.Accounts {
		firstMaskedPan := acc.MaskedPan[0]
		l4d := firstMaskedPan[len(firstMaskedPan)-4:]

		protoAccs[i] = &pbc.Account{
			Source:       acc.Source,
			ExternalId:   acc.ExternalId,
			CurrencyCode: acc.CurrencyCode,
			Balance:      acc.Balance,
			L4D:          l4d,
			Name:         acc.Name,
		}
	}

	protoJars := make([]*pbc.Jar, len(client.Jars))
	for i, jar := range client.Jars {
		protoJars[i] = &pbc.Jar{
			ExternalId:   jar.ExternalId,
			Title:        jar.Title,
			Description:  jar.Description,
			CurrencyCode: jar.CurrencyCode,
			Balance:      jar.Balance,
			Goal:         jar.Goal,
		}
	}

	msg := &pbc.Client{
		Source:     client.Source,
		ExternalId: client.ExternalId,
		Name:       client.Name,
		Accounts:   protoAccs,
		Jars:       protoJars,
	}

	return msg
}
