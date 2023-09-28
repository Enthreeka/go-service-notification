package entity

import "time"

type Notification struct {
	ID             string         `json:"id"`
	Message        string         `json:"message"`
	ClientProperty ClientProperty `json:"client_properties"`
	CreateAt       time.Time      `json:"create_at"`
	ExpiresAt      time.Time      `json:"expires_at"`
}
