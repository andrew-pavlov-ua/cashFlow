package handler

import (
	utils "github.com/andrew-pavlov-ua/pkg"
	"github.com/andrew-pavlov-ua/pkg/amqp"
	"github.com/andrew-pavlov-ua/pkg/logger"
	"github.com/andrew-pavlov-ua/pkg/models"
	"github.com/andrew-pavlov-ua/services/api-gateway/internal/clients"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	RabbitMQPublisher *amqp.RabbitMQPublisher
	Service           *clients.MonoClient
}

func NewHandler(amqpDSN string) *Handler {
	publisher, err := amqp.NewRabbitMQPublisher(amqpDSN)
	if err != nil {
		logger.Panic("Failed to initialize RabbitMQ publisher", err)
	}

	return &Handler{
		RabbitMQPublisher: publisher,
		Service:           clients.NewMonoClient(),
	}
}

func (h *Handler) NewTransaction(c *gin.Context) {
	var req models.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON request", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	msg := utils.TransactionRequestToProto(&req)

	// Publish the transaction to RabbitMQ
	err := h.RabbitMQPublisher.Publish(amqp.EXCHANGE, msg)
	if err != nil {
		logger.Error("Failed to publish transaction", err)
		c.JSON(500, gin.H{"error": "Failed to process transaction"})
		return
	}

	logger.Info("Transaction processed successfully")
	c.JSON(200, gin.H{"status": "Transaction processed successfully"})
}

func (h *Handler) InitClient(mono_token string) error {
	clientData, err := h.Service.FetchClient(mono_token)
	if err != nil {
		return err
	}
	clientData.Source = models.MONOBANK_SOURCE

	h.RabbitMQPublisher.Publish(amqp.ClientRoutingKey, utils.ClientToProto(clientData))

	logger.Infof("Initializing client with mono_token: %s", mono_token)
	return nil
}
