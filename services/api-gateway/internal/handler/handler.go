package handler

import (
	utils "github.com/andrew-pavlov-ua/pkg"
	"github.com/andrew-pavlov-ua/pkg/amqp"
	"github.com/andrew-pavlov-ua/pkg/logger"
	"github.com/andrew-pavlov-ua/pkg/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	RabbitMQPublisher *amqp.RabbitMQPublisher
}

func NewHandler(amqpDSN string) *Handler {
	publisher, err := amqp.NewRabbitMQPublisher(amqpDSN)
	if err != nil {
		logger.Panic("Failed to initialize RabbitMQ publisher", err)
	}

	return &Handler{
		RabbitMQPublisher: publisher,
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
	err := h.RabbitMQPublisher.Publish(amqp.Exchange, msg)
	if err != nil {
		logger.Error("Failed to publish transaction", err)
		c.JSON(500, gin.H{"error": "Failed to process transaction"})
		return
	}

	logger.Info("Transaction processed successfully")
	c.JSON(200, gin.H{"status": "Transaction processed successfully"})
}
