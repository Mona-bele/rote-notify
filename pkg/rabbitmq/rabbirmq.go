package rabbitmq

import (
	"fmt"

	"github.com/Mona-bele/logutils-go/logutils"
	"github.com/Mona-bele/rote-notify/pkg/env"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeName          = "ex_notifications_user_id"
	exchangeType          = "topic"
	TtlAmpqExpired365Days = int32(1471228928)
)

// RabbitMQ struct
type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

// Message struct
type Message struct {
	Type       string `json:"type"`
	UserID     string `json:"user_id"`
	RoutingKey string `json:"routing_key"`
	Body       []byte `json:"body"`
}

// CloseRabbitMQ closes the RabbitMQ connection
func (r *RabbitMQ) CloseRabbitMQ() {
	err := r.Ch.Close()
	if err != nil {
		logutils.Error("Failed to close the channel", err, nil)
	}

	err = r.Conn.Close()
	if err != nil {
		logutils.Error("Failed to close the connection", err, nil)
	}
	logutils.Info("RabbitMQ connection closed", nil)
}

// NewRabbitMQ creates a new RabbitMQ instance
func NewRabbitMQ(env *env.Env) *RabbitMQ {
	conn, ch := connectRabbitMQ(env)
	return &RabbitMQ{Conn: conn, Ch: ch}
}

// connectRabbitMQ to RabbitMQ
func connectRabbitMQ(env *env.Env) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(env.RabbitmqUrl)
	if err != nil {
		logutils.Error("Failed to connect to RabbitMQ", err, nil)
	}

	ch, err := conn.Channel()
	if err != nil {
		logutils.Error("Failed to open a channel", err, nil)
	}

	err = ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		logutils.Error("Failed to declare an exchange", err, nil)
	}

	logutils.Info("Connected to RabbitMQ", nil)

	return conn, ch
}

// CreateUserQueue Create a user-specific queue
func (r *RabbitMQ) CreateUserQueue(userID string, temporary bool) {
	queueName := "user_" + userID

	args := make(amqp.Table)
	if temporary {
		args["x-expires"] = TtlAmpqExpired365Days
	}
	args["x-message-ttl"] = TtlAmpqExpired365Days

	q, err := r.Ch.QueueDeclare(queueName, !temporary, false, false, false, args)
	if err != nil {
		logutils.Error("Failed to declare a queue", err, nil)
	}

	err = r.Ch.QueueBind(q.Name, fmt.Sprintf("user.%s.*", userID), exchangeName, false, nil)
	if err != nil {
		logutils.Error("Failed to bind a queue", err, nil)
	}

	args = nil

	logutils.Info("Queue created", map[string]interface{}{"queue": q.Name})
}

// DeleteUserQueue Delete a user-specific queue
func (r *RabbitMQ) DeleteUserQueue(userID string) {
	queueName := "user_" + userID
	_, err := r.Ch.QueueDelete(queueName, false, false, false)
	if err != nil {
		logutils.Error("Failed to delete a queue", err, nil)
	}
	logutils.Info("Queue deleted", map[string]interface{}{"queue": queueName})
}

// PublishMessage Publish a message to the exchange
func (r *RabbitMQ) PublishMessage(message Message) error {
	err := r.Ch.Publish(exchangeName, message.RoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message.Body,
	})
	if err != nil {
		logutils.Error("Failed to publish a message", err, nil)
		return err
	}
	logutils.Info("Message published", map[string]interface{}{"routing_key": message.RoutingKey})

	return nil
}

// ConsumeMessages Consume messages from the exchange
func (r *RabbitMQ) ConsumeMessages(userID string) <-chan amqp.Delivery {
	queueName := "user_" + userID
	msgs, err := r.Ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		logutils.Error("Failed to consume messages", err, nil)
	}
	logutils.Info("Consuming messages", map[string]interface{}{"queue": queueName})

	return msgs
}
