package notifications_user_id

import (
	"context"
	"encoding/json"
	"github.com/Mona-bele/commons-tools/pkg/adapter/rabbitmq"
	"github.com/Mona-bele/logutils-go/logutils"
	"github.com/Mona-bele/rote-notify/core/entity"
)

// NotificationsUserId struct
type NotificationsUserId struct {
	RabbitMQ rabbitmq.RabbitInterface
}

type Body struct {
	DeviceToken string `json:"device_token"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BodyPicle struct {
	RecipientID string `json:"recipient_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Type        string `json:"type"`
	IsRead      bool   `json:"is_read"`
	Token       string `json:"token,omitempty"`
	Args        []any  `json:"args,omitempty"`
}

func (b *Body) String() string {
	bodyJson, err := json.Marshal(b)
	if err != nil {
		logutils.Error("Failed to marshal the body", err, nil)
		return ""
	}

	return string(bodyJson)
}

// NewNotificationsUserId creates a new NotificationsUserId instance
func NewNotificationsUserId(url string) *NotificationsUserId {

	rmq := rabbitmq.NewRabbitmq(url, map[string]interface{}{})

	if err := rmq.ExchangeDeclare(rabbitmq.Exchange{
		Name:       "ex.picle.notification",
		Kind:       "topic",
		Durable:    true,
		AutoDelete: false,
	}); err != nil {
		logutils.Error("Failed to declare exchange", err, nil)
		return nil
	}

	if err := rmq.QueueDeclare(rabbitmq.Queue{
		Name:       "queue_picle_notification",
		Durable:    true,
		AutoDelete: false,
		Binds: &[]rabbitmq.Bind{
			{ExchangeName: "ex.picle.notification", BindingKey: "rk.picle.notification"},
			{ExchangeName: "ex.picle.notification", BindingKey: "rk.picle.notification.*"},
			{ExchangeName: "ex.picle.notification", BindingKey: "rk.picle.notification.app"},
			{ExchangeName: "ex.picle.notification", BindingKey: "rk.picle.notification.websocket"},
			{ExchangeName: "ex.picle.notification", BindingKey: "rk.picle.notification.email"},
		},
	}); err != nil {
		logutils.Error("Failed to declare queue rk.picle.notification.app", err, nil)
		return nil
	}

	return &NotificationsUserId{
		RabbitMQ: rmq,
	}
}

// NotifyPicle notifies the user ID
func (n *NotificationsUserId) NotifyPicle(ctx context.Context, userID string, body []byte, typeMessage entity.NotifyTypeMessage, args ...any) {
	if body == nil {
		body = []byte(typeMessage.GetNotifyTypeMessage())
	}

	bodyPicle := BodyPicle{
		RecipientID: userID,
		Title:       typeMessage.String(),
		Body:        string(body),
		Type:        typeMessage.String(),
		IsRead:      false,
		Args:        args,
	}

	bodyPicleJson, err := json.Marshal(bodyPicle)
	if err != nil {
		logutils.Error("Failed to marshal the body", err, nil)
		return
	}

	err = n.RabbitMQ.Producer(ctx, &rabbitmq.ProducerConfig{
		Exchange: "ex.picle.notification",
		Key:      "rk.picle.notification.app",
	},
		&rabbitmq.Message{
			Data:        bodyPicleJson,
			ContentType: "application/json",
		},
	)

	if err != nil {
		logutils.Error("Failed to publish to rk.picle.notification.app", err, nil)
		return
	}

	err = n.RabbitMQ.Producer(ctx, &rabbitmq.ProducerConfig{
		Exchange: "ex.picle.notification",
		Key:      "rk.picle.notification.websocket",
	},
		&rabbitmq.Message{
			Data:        bodyPicleJson,
			ContentType: "application/json",
		})

	if err != nil {
		logutils.Error("Failed to publish to rk.picle.notification.websocket", err, nil)
		return
	}
	logutils.Info("User ID notified", logutils.Fields{
		"user_id": userID,
		"type":    typeMessage.GetNotifyTypeMessage(),
	})
}

func (n *NotificationsUserId) NotifyApp(ctx context.Context, routingKey, userID string, body []byte, typeMessage string, args ...any) {
	if body == nil {
		body = []byte(typeMessage)
	}

	bodyPicle := BodyPicle{
		RecipientID: userID,
		Title:       typeMessage,
		Type:        typeMessage,
		Body:        string(body),
		IsRead:      false,
		Args:        args,
	}

	bodyPicleJson, err := json.Marshal(bodyPicle)
	if err != nil {
		logutils.Error("Failed to marshal the body", err, nil)
		return
	}

	if routingKey == "" {
		routingKey = "app"
	}

	err = n.RabbitMQ.Producer(ctx, &rabbitmq.ProducerConfig{
		Exchange: "ex.picle.notification",
		Key:      "rk.picle.notification." + routingKey,
	},
		&rabbitmq.Message{
			Data:        bodyPicleJson,
			ContentType: "application/json",
		})
	if err != nil {
		logutils.Error("Failed to publish to rk.picle.notification", err, nil)
		return
	}

	logutils.Info("User ID notified", logutils.Fields{
		"user_id": userID,
		"type":    typeMessage,
	})
}

// Close closes the RabbitMQ connection
func (n *NotificationsUserId) Close() error {
	if n.RabbitMQ != nil {
		return n.RabbitMQ.Close()
	}
	return nil
}
