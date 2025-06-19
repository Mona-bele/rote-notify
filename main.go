package main

import (
	"context"
	"github.com/joho/godotenv"
	"os"

	"github.com/Mona-bele/rote-notify/core/notifications_user_id"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	notifyService := notifications_user_id.NewNotificationsUserId(os.Getenv("RABBITMQ_URL"))

	//notifyService.RabbitMQ.CreateUserQueue("82181483-e936-4932-881f-bf7e28d06e5a", false)
	//body := []byte("account exists")
	notifyService.NotifyPicle(context.Background(), "d724bae2-8d0b-4145-9c41-6959735a6a9e", nil, "request_accepted", "Isac", 40)
	//time.Sleep(4 * time.Second)
	//notifyService.NotifyPicle(context.Background(), "85fc0acf-ca17-4e06-8cd3-0a1b968b6c86", []byte("Your balance is insufficient"), "balance_insufficient")

	//notifyService.DeleteNotificationsUserId(context.Background(), "82181483-e936-4932-881f-bf7e28d06e5a")

	//defer notifyService.CloseNotificationsUserId()
}
