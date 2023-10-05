package dto

import (
	"github.com/Enthreeka/go-service-notification/internal/entity"
)

type TimeNotificationRequest struct {
	CreateAt string `json:"create_at"`
}

type UpdateNotificationRequest struct {
	Message   string `json:"message"`
	CreateAt  string `json:"create_at"`
	ExpiresAt string `json:"expires_at"`
}

type CreateClientRequest struct {
	PhoneNumber    string                `json:"phone_number"`
	ClientProperty entity.ClientProperty `json:"client_property"`
	TimeZone       string                `json:"time_zone"`
}

type UpdateClientRequest struct {
	ID               string                `json:"id"`
	ClientPropertyID string                `json:"id_client_properties"`
	PhoneNumber      string                `json:"phone_number"`
	TimeZone         string                `json:"time_zone"`
	ClientProperty   entity.ClientProperty `json:"client_property"`
}

type CreateNotificationRequest struct {
	Message       string         `json:"message"`
	CreateAt      string         `json:"create_at"`
	ExpiresAt     string         `json:"expires_at"`
	Tags          []Tag          `json:"tags"`
	OperatorCodes []OperatorCode `json:"operator_codes"`
}

type Tag struct {
	Tag string `json:"tag"`
}

type OperatorCode struct {
	OperatorCode string `json:"operator_code"`
}
