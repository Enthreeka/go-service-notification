package entity

import (
	"strings"
	"time"
	"unicode/utf8"
)

type Client struct {
	ID               string         `json:"id"`
	ClientPropertyID string         `json:"id_client_properties"`
	PhoneNumber      string         `json:"phone_number"`
	ClientProperty   ClientProperty `json:"client_property"`
	TimeZone         time.Time      `json:"time_zone"`
}

type ClientProperty struct {
	Tag          string `json:"tag"`
	OperatorCode string `json:"operator_code"`
}

type Tag struct {
	Tag string `json:"tag"`
}

type OperatorCode struct {
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
