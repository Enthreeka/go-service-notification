package entity

import (
	"strings"
	"unicode/utf8"
)

type Client struct {
	ID           string `json:"id"`
	Tag          string `json:"tag"`
	PhoneNumber  string `json:"phone_number"`
	OperatorCode string `json:"operator_code"`
	TimeZone     string `json:"time_zone"`
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
