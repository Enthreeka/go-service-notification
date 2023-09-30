package entity

import "time"

type Notification struct {
	ID             string           `json:"id"`
	Message        string           `json:"message"`
	CreateAt       string           `json:"create_at"`
	ExpiresAt      string           `json:"expires_at"`
	ClientProperty []ClientProperty `json:"client_property"`
}

type ClientProperty struct {
	PhoneNumber  string `json:"phone_number"`
	OperatorCode string `json:"operator_code"`
}

func IsCorrectTime(value string) bool {
	callTime, err := time.ParseInLocation("15:04 02.01.2006", value, time.Local)
	if err != nil {
		return false
	}

	if !callTime.After(time.Now()) {
		return false
	}

	return true
}
