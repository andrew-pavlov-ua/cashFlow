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
	MonoClient        *clients.MonoClient
}

func NewHandler(amqpDSN string) *Handler {
	publisher, err := amqp.NewRabbitMQPublisher(amqpDSN)
	if err != nil {
		logger.Panic("Failed to initialize RabbitMQ publisher", err)
	}

	return &Handler{
		RabbitMQPublisher: publisher,
		MonoClient:        clients.NewMonoClient(),
	}
}

func (h *Handler) WebhookHandle(c *gin.Context) {
	var req models.MonoWebhook
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON request", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	if req.Data.StatementItem.ExternalId == "" {
		req.Data.StatementItem.ExternalId = req.Data.Account
	}

	logger.Info(req.PrettyLog())

	msg := utils.TransactionRequestToProto(&req.Data.StatementItem)

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

func (h *Handler) NewClient(c *gin.Context) {
	var req struct {
		MonoToken string `json:"mono_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to bind JSON request", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := h.InitClient(req.MonoToken)
	if err != nil {
		logger.Error("Failed to initialize client", err)
		c.JSON(500, gin.H{"error": "Failed to initialize client"})
		return
	}

	logger.Info("Client initialized successfully")

	logger.Infof("Setting up MonoBank webhook for token: %s", req.MonoToken)
	err = h.MonoClient.SetUpMonoWebhook(req.MonoToken)
	if err != nil {
		logger.Error("Failed to set up MonoBank webhook", err)
		c.JSON(500, gin.H{"error": "Failed to set up MonoBank webhook"})
		return
	}
	logger.Info("MonoBank webhook set up successfully")

	c.JSON(200, gin.H{"status": "Client initialized successfully"})
}

func (h *Handler) InitClient(mono_token string) error {
	clientData, err := h.MonoClient.FetchClient(mono_token)
	if err != nil {
		return err
	}
	clientData.Source = models.MONOBANK_SOURCE

	h.RabbitMQPublisher.Publish(amqp.ClientRoutingKey, utils.ClientToProto(clientData))

	logger.Infof("Initializing client with mono_token: %s", mono_token)
	return nil
}

func (h *Handler) WebhookConfirmation(c *gin.Context) {
	// MonoBank webhook confirmation endpoint
	c.String(200, "OK")
}
