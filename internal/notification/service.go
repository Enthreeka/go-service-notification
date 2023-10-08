package notification

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
)

type ClientService interface {
	CreateClient(ctx context.Context, request *dto.CreateClientRequest) error
	UpdateClient(ctx context.Context, request *dto.UpdateClientRequest) error
	DeleteClient(ctx context.Context, id string) error
}

type NotificationService interface {
	CreateNotification(ctx context.Context, request *dto.CreateNotificationRequest) error
	UpdateNotification(ctx context.Context, request *dto.UpdateNotificationRequest) error
	DeleteNotification(ctx context.Context, request *dto.TimeNotificationRequest) error
	GetByCreateTime(ctx context.Context, request *dto.TimeNotificationRequest) ([]entity.Notification, error)
}

type MessageService interface {
	GetInfoNotification(ctx context.Context, id string) (map[string][]entity.MessageInfo, error)
	GetAllGroupByStatus(ctx context.Context) (map[string][]entity.MessageInfo, error)
}
