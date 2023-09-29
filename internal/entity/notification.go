package entity

import "time"

type Notification struct {
	ID           string    `json:"id"`
	Message      string    `json:"message"`
	PhoneNumber  string    `json:"phone_number"`
	OperatorCode string    `json:"operator_code"`
	CreateAt     time.Time `json:"create_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}
