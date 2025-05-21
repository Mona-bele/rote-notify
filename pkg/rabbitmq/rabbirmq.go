package rabbitmq

import (
	"context"
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
	conn        *amqp.Connection
	ch          *amqp.Channel
	env         *env.Env
	reconnectCh chan struct{}
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
	rmq := &RabbitMQ{
		env:         env,
		reconnectCh: make(chan struct{}),
	}
	go rmq.connectWithRetry()
	return rmq
}

func (r *RabbitMQ) connectWithRetry() {
	for {
		conn, err := amqp.Dial(r.env.RabbitmqUrl)
		if err != nil {
			logutils.Error("Failed to connect to RabbitMQ, retrying...", err, nil)
			time.Sleep(5 * time.Second)
			continue
		}

		ch, err := conn.Channel()
		if err != nil {
			logutils.Error("Failed to open a channel, retrying...", err, nil)
			_ = conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		err = ch.Confirm(false)
		if err != nil {
			logutils.Error("Failed to put channel in confirm mode", err, nil)
		}

		r.conn = conn
		r.ch = ch

		err = ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
		if err != nil {
			logutils.Error("Failed to declare an exchange", err, nil)
		}

		logutils.Info("Connected to RabbitMQ and channel ready", nil)

		notifyClose := make(chan *amqp.Error)
		r.conn.NotifyClose(notifyClose)

		select {
		case err := <-notifyClose:
			logutils.Warn("RabbitMQ connection closed, reconnecting...", logutils.Fields{"error": err})
			time.Sleep(2 * time.Second)
		}
	}
}

func (r *RabbitMQ) CloseRabbitMQ() {
	if r.ch != nil {
		_ = r.ch.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
	logutils.Info("RabbitMQ connection closed", nil)
}

// PublishMessage Publish a message to the exchange
func (r *RabbitMQ) PublishMessage(message Message, contentType string) error {
	for i := 0; i < 3; i++ {
		err := r.ch.PublishWithContext(context.Background(),
			exchangeName, message.RoutingKey, false, false, amqp.Publishing{
				ContentType: contentType,
				Body:        message.Body,
			})
		if err != nil {
			logutils.Error(fmt.Sprintf("Publish attempt %d failed", i+1), err, nil)
			time.Sleep(time.Duration(2*i+1) * time.Second)
			continue
		}

		select {
		case confirm := <-r.ch.NotifyPublish(make(chan amqp.Confirmation, 1)):
			if confirm.Ack {
				logutils.Info("Message confirmed", map[string]interface{}{"routing_key": message.RoutingKey})
				return nil
			}
			logutils.Warn("Message not acknowledged, retrying...", nil)
		case <-time.After(5 * time.Second):
			logutils.Warn("No confirm received in time, retrying...", nil)
		}
	}
	return fmt.Errorf("failed to publish message after retries")
}

func (r *RabbitMQ) WaitForReady(timeout time.Duration) error {
	start := time.Now()
	for {
		if r.ch != nil && !r.ch.IsClosed() {
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

	q, err := r.ch.QueueDeclare("queue_picle_notification", true, false, false, false, nil)
	if err != nil {
		logutils.Error("Failed to declare queue_picle_notification", err, nil)
		return
	}

	err = r.ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	if err != nil {
		logutils.Error("Failed to declare an exchange", err, nil)
	}

	err = r.ch.QueueBind(q.Name, "rk.picle.notification", exchangeName, false, nil)
	if err != nil {
		logutils.Error("Failed to bind queue_picle_notification", err, nil)
		return
	}

	err = r.ch.QueueBind(q.Name, "rk.picle.notification.email", exchangeName, false, nil)
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
