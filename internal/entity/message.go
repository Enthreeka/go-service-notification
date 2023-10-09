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
	ID             string `json:"id"`
	NotificationID string `json:"notification_id"`
	CreatedAt      string `json:"created_at"`
	Message        string `json:"message"`
	ExpiresAt      string `json:"expires_at"`
	ClientID       string `json:"client_id"`
	PhoneNumber    string `json:"phone_number"`

	Status string `json:"-"`
	InTime bool   `json:"-"`
}

type MessageInfo struct {
	CreatedAt    string `json:"created_at,omitempty"`
	Status       string `json:"status,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	Tag          string `json:"tag,omitempty"`
	OperatorCode string `json:"operator_code,omitempty"`
	Message      string `json:"message,omitempty"`

	Count          int    `json:"count,omitempty"`
	NotificationID string `json:"notification_id,omitempty"`
}
