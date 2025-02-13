package notifications_user_id

import (
	"context"
	"encoding/json"
	"github.com/Mona-bele/logutils-go/logutils"
	"github.com/Mona-bele/rote-notify/core/entity"
	"github.com/Mona-bele/rote-notify/pkg/env"
	"github.com/Mona-bele/rote-notify/pkg/rabbitmq"
	"github.com/Mona-bele/rote-notify/pkg/security/jwt"
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
	// Convert the Body struct to a JSON string
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

/*// NotifyUserId notifies the user ID
func (n *NotificationsUserId) NotifyUserId(ctx context.Context, userID string, typeMessage entity.NotifyTypeMessage) {

	n.RabbitMQ.CreateUserQueue(userID, false)

	body := Body{
		Title:       typeMessage.String(),
		Description: typeMessage.GetNotifyTypeMessage(),
	}

	token, err := n.jwt.GenerateToken(body.String(), n.env.JwtIssuer, n.env.JwtAudience, n.env.JwtSubject)
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

	err = n.RabbitMQ.PublishMessage(message, "text/plain")
	if err != nil {
		logutils.Error("Failed to publish a message", err, nil)
		return
	}

class NotificationInitializer extends ConsumerWidget {
  const NotificationInitializer({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    // Inicializar o serviço de notificações passando o ref
    PushNotificationService.initializeNotificationService(ref);

    return const AppWidget();
  }
}

	// Queue picle
	n.NotifyPicle(ctx, userID, typeMessage)

	logutils.Info("User ID notified", logutils.Fields{"user_id": userID, "type": typeMessage.GetNotifyTypeMessage()})
}*/

// NotifyPicle notifies the user ID
func (n *NotificationsUserId) NotifyPicle(ctx context.Context, userID string, body []byte, typeMessage entity.NotifyTypeMessage) {

	n.RabbitMQ.QueuePicle()

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
	}

	messagePicle := rabbitmq.Message{
		Type:       typeMessage.String(),
		UserID:     userID,
		RoutingKey: "rk.picle.notification",
		Body:       bodyPicleJson,
	}

	err = n.RabbitMQ.PublishMessage(messagePicle, "application/json")
	if err != nil {
		logutils.Error("Failed to publish a message to picle", err, nil)
		return
	}

	logutils.Info("User ID notified", logutils.Fields{"user_id": userID, "type": typeMessage.GetNotifyTypeMessage()})
}

/*// DeleteNotificationsUserId deletes the user ID if exists messages in the queue
func (n *NotificationsUserId) DeleteNotificationsUserId(ctx context.Context, userID string) {
	msgs, err := n.RabbitMQ.VerifyMessageInQueue(userID)
	if err != nil {
		logutils.Error("Failed to verify messages in the queue", err, nil)
		return
	}

	logutils.Warn("User ID has messages in the queue", logutils.Fields{"user_id": userID, "messages": msgs})
	if msgs < 0 {
		logutils.Info("User ID deleted", logutils.Fields{"user_id": userID})
		n.RabbitMQ.DeleteUserQueue(userID)
	}

}*/

// CloseNotificationsUserId closes the RabbitMQ connection
func (n *NotificationsUserId) CloseNotificationsUserId() {
	n.RabbitMQ.CloseRabbitMQ()
}
