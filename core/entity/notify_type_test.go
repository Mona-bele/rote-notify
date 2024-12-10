package entity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNotifyTypeMessage(t *testing.T) {
	// TestGetNotifyTypeMessage tests the GetNotifyTypeMessage function
	// It should return a string message
	// Arrange
	type args struct {
		t NotifyTypeMessage
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test GetNotifyTypeMessage",
			args: args{
				t: DEPOSIT,
			},
			want: "Deposit completed",
		},
		{
			name: "Test GetNotifyTypeMessage",
			args: args{
				t: WITHDRAW,
			},
			want: "Withdraw completed",
		},
		{
			name: "Test GetNotifyTypeMessage",
			args: args{
				t: TRANSFER,
			},
			want: "Transfer completed",
		},
		{
			name: "Test GetNotifyTypeMessage",
			args: args{
				t: REQUEST_EXCHANGE,
			},
			want: "New purchase request",
		},
		{
			name: "Test GetNotifyTypeMessage",
			args: args{
				t: REQUEST_EXPIRED,
			},
			want: "Expired purchase request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNotifyTypeMessage(tt.args.t); got != tt.want {
				t.Errorf("GetNotifyTypeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyTypeMessage_GetNotifyTypeMessage(t *testing.T) {
	// TestNotifyTypeMessage_GetNotifyTypeMessage tests the GetNotifyTypeMessage method
	// It should return a string message
	// Arrange
	tests := []struct {
		name string
		t    NotifyTypeMessage
		want string
	}{
		{
			name: "Test GetNotifyTypeMessage",
			t:    DEPOSIT,
			want: "Deposit completed",
		},
		{
			name: "Test GetNotifyTypeMessage",
			t:    WITHDRAW,
			want: "Withdraw completed",
		},
		{
			name: "Test GetNotifyTypeMessage",
			t:    TRANSFER,
			want: "Transfer completed",
		},
		{
			name: "Test GetNotifyTypeMessage",
			t:    REQUEST_EXCHANGE,
			want: "New purchase request",
		},
		{
			name: "Test GetNotifyTypeMessage",
			t:    REQUEST_EXPIRED,
			want: "Expired purchase request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.GetNotifyTypeMessage()
			if got != tt.want {
				t.Errorf("NotifyTypeMessage.GetNotifyTypeMessage() = %v, want %v", got, tt.want)
			}
			fmt.Printf("NotifyTypeMessage.GetNotifyTypeMessage() = %v", got)
			assert.Equal(t, tt.want, got, "NotifyTypeMessage.GetNotifyTypeMessage() = %v, want %v", got, tt.want)
		})
	}
}
