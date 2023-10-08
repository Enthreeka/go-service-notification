package entity

import "time"

type Message struct {
	ID             string    `json:"id"`
	NotificationID string    `json:"notification_id"`
	ClientID       string    `json:"client_id"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type ClientsMessage struct {
	NotificationID string `json:"notification_id"`
	CreatedAt      string `json:"created_at"`
	Message        string `json:"message"`
	ExpiresAt      string `json:"expires_at"`
	ClientID       string `json:"client_id"`
	PhoneNumber    string `json:"phone_number"`

	Status string `json:"-"`
}
