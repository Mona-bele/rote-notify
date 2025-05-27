package main

import (
	"context"
	"github.com/Mona-bele/rote-notify/core/entity"

	"github.com/Mona-bele/rote-notify/core/notifications_user_id"
	"github.com/Mona-bele/rote-notify/pkg/env"
)

func main() {
	notifyService := notifications_user_id.NewNotificationsUserId(env.LoadEnv(".env"))

	//notifyService.RabbitMQ.CreateUserQueue("82181483-e936-4932-881f-bf7e28d06e5a", false)
	//body := []byte("account exists")
	notifyService.NotifyPicle(context.Background(), "bd4a2d2f-9812-48b8-8352-011578aa0153", nil, entity.DEPOSIT_PROCESS)
	//time.Sleep(4 * time.Second)
	//notifyService.NotifyPicle(context.Background(), "85fc0acf-ca17-4e06-8cd3-0a1b968b6c86", []byte("Your balance is insufficient"), "balance_insufficient")

	//notifyService.DeleteNotificationsUserId(context.Background(), "82181483-e936-4932-881f-bf7e28d06e5a")

	//defer notifyService.CloseNotificationsUserId()
}
