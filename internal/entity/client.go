package entity

import (
	"strings"
	"time"
	"unicode/utf8"
)

type Client struct {
	ID             string         `json:"id"`
	Tag            string         `json:"tag"`
	ClientProperty ClientProperty `json:"client_properties"`
	TimeZone       time.Time      `json:"time_zone"`
}

type ClientProperty struct {
	ID           string `json:"id"`
	PhoneNumber  string `json:"phone_number"`
	OperatorCode string `json:"operator_code"`
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
