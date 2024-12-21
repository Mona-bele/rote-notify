package entity

// NotifyType struct
type NotifyType struct {
	Type   string `json:"type"`
	UserID string `json:"user_id"`
	Body   []byte `json:"body"`
}

type NotifyTypeMessage string

const (
	// Deposit
	DEPOSIT         NotifyTypeMessage = NotifyTypeMessage("deposit")
	DEPOSIT_ERROR   NotifyTypeMessage = NotifyTypeMessage("deposit_error")
	DEPOSIT_SUCCESS NotifyTypeMessage = NotifyTypeMessage("deposit_success")
	DEPOSIT_CANCEL  NotifyTypeMessage = NotifyTypeMessage("deposit_cancel")
	DEPOSIT_PROCESS NotifyTypeMessage = NotifyTypeMessage("deposit_process")

	// Withdraw
	WITHDRAW         NotifyTypeMessage = NotifyTypeMessage("withdraw")
	WITHDRAW_ERROR   NotifyTypeMessage = NotifyTypeMessage("withdraw_error")
	WITHDRAW_SUCCESS NotifyTypeMessage = NotifyTypeMessage("withdraw_success")
	WITHDRAW_CANCEL  NotifyTypeMessage = NotifyTypeMessage("withdraw_cancel")
	WITHDRAW_PROCESS NotifyTypeMessage = NotifyTypeMessage("withdraw_process")

	// Transfer
	TRANSFER         NotifyTypeMessage = NotifyTypeMessage("transfer")
	TRANSFER_ERROR   NotifyTypeMessage = NotifyTypeMessage("transfer_error")
	TRANSFER_SUCCESS NotifyTypeMessage = NotifyTypeMessage("transfer_success")
	TRANSFER_CANCEL  NotifyTypeMessage = NotifyTypeMessage("transfer_cancel")
	TRANSFER_PROCESS NotifyTypeMessage = NotifyTypeMessage("transfer_process")

	// Request
	REQUEST_EXCHANGE  NotifyTypeMessage = NotifyTypeMessage("request_exchange")
	REQUEST_EXPIRED   NotifyTypeMessage = NotifyTypeMessage("request_expired")
	REQUEST_ACCEPTED  NotifyTypeMessage = NotifyTypeMessage("request_accepted")
	REQUEST_REJECTED  NotifyTypeMessage = NotifyTypeMessage("request_rejected")
	REQUEST_COMPLETED NotifyTypeMessage = NotifyTypeMessage("request_completed")
	REQUEST_PROCESS   NotifyTypeMessage = NotifyTypeMessage("request_process")
	REQUEST_CANCEL    NotifyTypeMessage = NotifyTypeMessage("request_cancel")

	// Post
	NEW_POST NotifyTypeMessage = NotifyTypeMessage("new_post")
)

// MapNotifyTypeMessage maps the NotifyTypeMessage to a string messages
var MapNotifyTypeMessage = map[NotifyTypeMessage]string{
	// Deposit
	DEPOSIT:         "Deposit completed",
	DEPOSIT_ERROR:   "Error occurred while processing the deposit",
	DEPOSIT_SUCCESS: "Deposit completed",
	DEPOSIT_CANCEL:  "Deposit canceled",
	DEPOSIT_PROCESS: "Processing the deposit",

	// Withdraw
	WITHDRAW:         "Withdraw completed",
	WITHDRAW_ERROR:   "Error occurred while processing the withdraw",
	WITHDRAW_SUCCESS: "Withdraw completed",
	WITHDRAW_CANCEL:  "Withdraw canceled",
	WITHDRAW_PROCESS: "Processing the withdraw",

	// Transfer
	TRANSFER:         "Transfer completed",
	TRANSFER_ERROR:   "Error occurred while processing the transfer",
	TRANSFER_SUCCESS: "Transfer completed",
	TRANSFER_CANCEL:  "Transfer canceled",
	TRANSFER_PROCESS: "Processing the transfer",

	// Request
	REQUEST_EXCHANGE:  "New purchase request",
	REQUEST_EXPIRED:   "Expired purchase request",
	REQUEST_ACCEPTED:  "Purchase request accepted",
	REQUEST_REJECTED:  "Purchase request rejected",
	REQUEST_COMPLETED: "Order completed",
	REQUEST_PROCESS:   "Processing the request",
	REQUEST_CANCEL:    "An error occurred while processing the request",

	// Post
	NEW_POST: "New post",
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
