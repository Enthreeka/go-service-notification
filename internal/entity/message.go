package entity

import "time"

type Message struct {
	ID             string    `json:"id"`
	NotificationID string    `json:"notification_id"`
	ClientID       string    `json:"client_id"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}
