package rabbitmq

import (
	"fmt"
	"sync"
	"time"

	"github.com/Mona-bele/logutils-go/logutils"
	"github.com/Mona-bele/rote-notify/pkg/env"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	exchangeName          = "ex.picle.notification"
	exchangeType          = "topic"
	TtlAmpqExpired365Days = int32(1471228928)
)

// RabbitMQ struct
type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
	env  *env.Env
}

// Message struct
type Message struct {
	Type       string `json:"type"`
	UserID     string `json:"user_id"`
	RoutingKey string `json:"routing_key"`
	Body       []byte `json:"body"`
}

// NewRabbitMQ creates a new RabbitMQ instance
func NewRabbitMQ(env *env.Env) *RabbitMQ {
	conn, ch := connectRabbitMQ(env)
	return &RabbitMQ{Conn: conn, Ch: ch}
}

func connectRabbitMQ(env *env.Env) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(env.RabbitmqUrl)
	if err != nil {
		logutils.Error("Failed to connect to RabbitMQ", err, nil)
		panic("Cannot continue without RabbitMQ connection")
	}

	ch, err := conn.Channel()
	if err != nil {
		logutils.Error("Failed to open a channel", err, nil)
		panic("Cannot continue without a RabbitMQ channel")
	}

	err = ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		logutils.Error("Failed to declare an exchange", err, nil)
		panic("Cannot continue without exchange declaration")
	}

	logutils.Info("Connected to RabbitMQ", nil)
	return conn, ch
}

func (r *RabbitMQ) CloseRabbitMQ() {
	if r.Ch != nil {
		_ = r.Ch.Close()
	}
	if r.Conn != nil {
		_ = r.Conn.Close()
	}
	logutils.Info("RabbitMQ connection closed", nil)
}

// PublishMessage Publish a message to the exchange
func (r *RabbitMQ) PublishMessage(message Message, contentType string) error {
	var lastErr error
	for i := 0; i < 3; i++ { // tenta até 3 vezes
		err := r.Ch.Publish(exchangeName, message.RoutingKey, false, false, amqp.Publishing{
			ContentType: contentType,
			Body:        message.Body,
		})
		if err == nil {
			logutils.Info("Message published", map[string]interface{}{"routing_key": message.RoutingKey})
			return nil
		}

		lastErr = err
		logutils.Warn("Retrying to publish message", map[string]interface{}{"attempt": i + 1, "error": err.Error()})
		time.Sleep(time.Duration(i+1) * time.Second) // backoff linear
	}

	logutils.Error("Failed to publish a message after retries", lastErr, nil)
	return lastErr
}

func (r *RabbitMQ) WaitForReady(timeout time.Duration) error {
	start := time.Now()
	for {
		if r.Ch != nil && !r.Ch.IsClosed() {
			return nil
		}
		if time.Since(start) > timeout {
			return fmt.Errorf("RabbitMQ connection timeout after %s", timeout)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

var picleQueueDeclared bool
var picleQueueMutex sync.Mutex

func (r *RabbitMQ) QueuePicle() {
	picleQueueMutex.Lock()
	defer picleQueueMutex.Unlock()

	if picleQueueDeclared {
		return
	}

	q, err := r.Ch.QueueDeclare("queue_picle_notification", true, false, false, false, nil)
	if err != nil {
		logutils.Error("Failed to declare queue_picle_notification", err, nil)
		return
	}

	err = r.Ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		logutils.Error("Failed to declare an exchange", err, nil)
	}

	err = r.Ch.QueueBind(q.Name, "rk.picle.notification", exchangeName, false, nil)
	if err != nil {
		logutils.Error("Failed to bind queue_picle_notification", err, nil)
		return
	}

	err = r.Ch.QueueBind(q.Name, "rk.picle.notification.email", exchangeName, false, nil)
	if err != nil {
		logutils.Error("Failed to bind rk.picle.notification.email", err, nil)
		return
	}

	picleQueueDeclared = true
	logutils.Info("queue_picle_notification declared and bound", map[string]interface{}{"queue": q.Name})
}

func rejectRabbitMessage(msg amqp.Delivery) {
	if err := msg.Reject(true); err != nil {
		logutils.Error("Error rejecting message: ", err, nil)
	}
}
