package entity

// NotifyType struct
type NotifyType struct {
	Type   string `json:"type"`
	UserID string `json:"user_id"`
	Body   []byte `json:"body"`
}

type NotifyTypeMessage string

const (
	DEPOSIT          NotifyTypeMessage = NotifyTypeMessage("deposit")
	WITHDRAW         NotifyTypeMessage = NotifyTypeMessage("withdraw")
	TRANSFER         NotifyTypeMessage = NotifyTypeMessage("transfer")
	REQUEST_EXCHANGE NotifyTypeMessage = NotifyTypeMessage("request_exchange")
	REQUEST_EXPIRED  NotifyTypeMessage = NotifyTypeMessage("request_expired")
	REQUEST_ACCEPTED NotifyTypeMessage = NotifyTypeMessage("request_accepted")
	NEW_POST         NotifyTypeMessage = NotifyTypeMessage("new_post")
)

// MapNotifyTypeMessage maps the NotifyTypeMessage to a string messages
var MapNotifyTypeMessage = map[NotifyTypeMessage]string{
	DEPOSIT:          "Deposit completed",
	WITHDRAW:         "Withdraw completed",
	TRANSFER:         "Transfer completed",
	REQUEST_EXCHANGE: "New purchase request",
	REQUEST_EXPIRED:  "Expired purchase request",
	REQUEST_ACCEPTED: "Purchase request accepted",
	NEW_POST:         "New post",
}

// GetNotifyTypeMessage returns the message of the NotifyTypeMessage
func GetNotifyTypeMessage(t NotifyTypeMessage) string {
	return MapNotifyTypeMessage[t]
}

// GetNotifyTypeMessage returns the message of the NotifyTypeMessage
func (t NotifyTypeMessage) GetNotifyTypeMessage() string {
	return MapNotifyTypeMessage[t]
}

// GetKey returns the key of the NotifyTypeMessage
func (t NotifyTypeMessage) String() string {
	return string(t)
}
