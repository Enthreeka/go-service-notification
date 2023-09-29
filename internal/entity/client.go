package entity

import (
	"strings"
	"time"
	"unicode/utf8"
)

type Client struct {
	ID           string    `json:"id"`
	Tag          string    `json:"tag"`
	PhoneNumber  string    `json:"phone_number"`
	OperatorCode string    `json:"operator_code"`
	TimeZone     time.Time `json:"time_zone"`
	TimeZoneDTO  string    `json:"time_zone_dto"`
}

func IsCorrectNumber(phoneNumber string) bool {
	if !strings.HasPrefix(phoneNumber, "7") {
		return false
	}

	if utf8.RuneCountInString(phoneNumber) != 11 {
		return false
	}

	return true
}
