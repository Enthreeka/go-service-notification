package dto

type DeleteNotificationRequest struct {
	CreateAt string `json:"create_at"`
}

type UpdateNotificationRequest struct {
	Message   string `json:"message"`
	CreateAt  string `json:"create_at"`
	ExpiresAt string `json:"expires_at"`
}

type GetNotificationRequest struct {
	CreateAt string `json:"create_at"`
}
