package amqp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/andrew-pavlov-ua/pkg/logger"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

const Exchange = "transactions"

func NewRabbitMQPublisher(amqpDSN string) (*RabbitMQPublisher, error) {
	logger.Info("Connecting to RabbitMQ...")

	conn, err := amqp.Dial(amqpDSN)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	err = ch.ExchangeDeclare(
		Exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		logger.Panic("panic during declaring ampq Exchange", err)
	}

	logger.Info("--- RabbitMQ Exchanges declared successfully ---")

	return &RabbitMQPublisher{
		conn:    conn,
		channel: ch,
	}, nil
}

func (p *RabbitMQPublisher) Publish(routingKey string, msg proto.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := proto.Marshal(msg)
	if err != nil {
		logger.Error("Failed to marshal message: ", err)
		return err
	}

	err = p.channel.PublishWithContext(
		ctx,
		Exchange,   // Exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "application/x-protobuf",
			Timestamp:    time.Now(),
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)

	if err != nil {
		// Check specific error types
		if errors.Is(err, context.DeadlineExceeded) {
			logger.Errorw("Publish timeout - RabbitMQ not responding in time",
				"Exchange", Exchange,
				"routing_key", routingKey,
				"error", err,
			)
			return fmt.Errorf("publish timeout: %w", err)
		}

		if errors.Is(err, context.Canceled) {
			logger.Warnw("Publish cancelled",
				"Exchange", Exchange,
				"routing_key", routingKey,
			)
			return fmt.Errorf("publish cancelled: %w", err)
		}

		// Other errors (connection closed, etc)
		logger.Errorw("Failed to publish message",
			"Exchange", Exchange,
			"routing_key", routingKey,
			"error", err,
		)
		return fmt.Errorf("failed to publish: %w", err)
	}

	logger.Debugf("Published to Exchange=%s, routingKey=%s", Exchange, routingKey)
	return nil
}

func (p *RabbitMQPublisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
}
