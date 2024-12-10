package notifications_user_id

import (
	"fmt"
	"github.com/Mona-bele/logutils-go/logutils"
	"ithub.com/Mona-bele/rote-notify/core/entity"
	"ithub.com/Mona-bele/rote-notify/pkg/env"
	"ithub.com/Mona-bele/rote-notify/pkg/rabbitmq"
	"ithub.com/Mona-bele/rote-notify/pkg/security/jwt"
)

// NotificationsUserId struct
type NotificationsUserId struct {
	env      env.Env
	RabbitMQ *rabbitmq.RabbitMQ
	jwt      *jwt.JWT
}

// NewNotificationsUserId creates a new NotificationsUserId instance
func NewNotificationsUserId(env env.Env) *NotificationsUserId {

	rmq := rabbitmq.NewRabbitMQ(env)
	jwt, err := jwt.NewJWTFromEnv(env)
	if err != nil {
		logutils.Error("Failed to create a new JWT instance", err, nil)
		return nil
	}

	return &NotificationsUserId{
		env:      env,
		RabbitMQ: rmq,
		jwt:      jwt,
	}
}

// NotifyUserId notifies the user ID
func (n *NotificationsUserId) NotifyUserId(userID string, typeMessage entity.NotifyTypeMessage) {

	n.RabbitMQ.CreateUserQueue(userID, false)

	token, err := n.jwt.GenerateToken(typeMessage.GetNotifyTypeMessage(), n.env.JwtIssuer, n.env.JwtAudience, n.env.JwtSubject)
	if err != nil {
		logutils.Error("Failed to generate a JWT token", err, nil)
		return
	}

	message := rabbitmq.Message{
		Type:       typeMessage.String(),
		UserID:     userID,
		RoutingKey: fmt.Sprintf("user.%s.%s", userID, typeMessage.String()),
		Body:       []byte(token),
	}

	err = n.RabbitMQ.PublishMessage(message)
	if err != nil {
		logutils.Error("Failed to publish a message", err, nil)
		return
	}

	logutils.Info("User ID notified", logutils.Fields{"user_id": userID, "type": typeMessage.GetNotifyTypeMessage()})
}

// DeleteNotificationsUserId deletes the user ID
func (n *NotificationsUserId) DeleteNotificationsUserId(userID string) {
	n.RabbitMQ.DeleteUserQueue(userID)
}

// CloseNotificationsUserId closes the RabbitMQ connection
func (n *NotificationsUserId) CloseNotificationsUserId() {
	n.RabbitMQ.CloseRabbitMQ()
}
