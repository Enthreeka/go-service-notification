package entity

import (
	"fmt"
	"time"
)

type Notification struct {
	ID               string   `json:"id,omitempty"`
	Message          string   `json:"message"`
	CreateAt         string   `json:"create_at"`
	ExpiresAt        string   `json:"expires_at"`
	ClientPropertyID []string `json:"id_client_properties"`

	ClientProperty []ClientProperty `json:"client_property"`
}

func IsCorrectTime(value string) bool {
	callTime, err := time.ParseInLocation("15:04 02.01.2006", value, time.Local)
	if err != nil {
		return false
	}

	if !callTime.After(time.Now()) {
		fmt.Println("Test")
		return false
	}

	return true
}
