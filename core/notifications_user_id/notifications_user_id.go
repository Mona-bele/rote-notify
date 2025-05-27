package notifications_user_id

import (
	"context"
	"encoding/json"
	"github.com/Mona-bele/logutils-go/logutils"
	"github.com/Mona-bele/rote-notify/core/entity"
	"github.com/Mona-bele/rote-notify/pkg/env"
	"github.com/Mona-bele/rote-notify/pkg/rabbitmq"
	"github.com/Mona-bele/rote-notify/pkg/security/jwt"
	"time"
)

// NotificationsUserId struct
type NotificationsUserId struct {
	env      *env.Env
	RabbitMQ *rabbitmq.RabbitMQ
	jwt      *jwt.JWT
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
func NewNotificationsUserId(env *env.Env) *NotificationsUserId {

	rmq := rabbitmq.NewRabbitMQ(env)

	if err := rmq.WaitForReady(10 * time.Second); err != nil {
		logutils.Error("RabbitMQ not ready in time", err, nil)
		return nil
	}

	rmq.QueuePicle()

	return &NotificationsUserId{
		env:      env,
		RabbitMQ: rmq,
	}
}

// NotifyPicle notifies the user ID
func (n *NotificationsUserId) NotifyPicle(ctx context.Context, userID string, body []byte, typeMessage entity.NotifyTypeMessage) {
	if body == nil {
		body = []byte(typeMessage.GetNotifyTypeMessage())
	}

	bodyPicle := BodyPicle{
		RecipientID: userID,
		Title:       typeMessage.String(),
		Body:        string(body),
		Type:        typeMessage.String(),
		IsRead:      false,
	}

	bodyPicleJson, err := json.Marshal(bodyPicle)
	if err != nil {
		logutils.Error("Failed to marshal the body", err, nil)
		return
	}

	// Envia para notificação em app
	err = n.RabbitMQ.PublishMessage(rabbitmq.Message{
		Type:       typeMessage.String(),
		UserID:     userID,
		RoutingKey: "rk.picle.notification.app",
		Body:       bodyPicleJson,
	}, "application/json")

	if err != nil {
		logutils.Error("Failed to publish to rk.picle.notification", err, nil)
		return
	}

	err = n.RabbitMQ.PublishMessage(rabbitmq.Message{
		Type:       typeMessage.String(),
		UserID:     userID,
		RoutingKey: "rk.picle.notification.websocket",
		Body:       bodyPicleJson,
	}, "application/json")

	if err != nil {
		logutils.Error("Failed to publish to rk.picle.notification.email", err, nil)
		return
	}
	logutils.Info("User ID notified", logutils.Fields{
		"user_id": userID,
		"type":    typeMessage.GetNotifyTypeMessage(),
	});
}

/*
	// Envia para notificação via e-mail
	err = n.RabbitMQ.PublishMessage(rabbitmq.Message{
		Type:       typeMessage.String(),
		UserID:     userID,
		RoutingKey: "rk.picle.notification.email",
		Body:       bodyPicleJson,
	}, "application/json")

	if err != nil {
		logutils.Error("Failed to publish to rk.picle.notification.email", err, nil)
		return
	}

	logutils.Info("User ID notified", logutils.Fields{
		"user_id": userID,
		"type":    typeMessage.GetNotifyTypeMessage(),
	})
*/

// CloseNotificationsUserId closes the RabbitMQ connection
func (n *NotificationsUserId) CloseNotificationsUserId() {
	n.RabbitMQ.CloseRabbitMQ()
}
