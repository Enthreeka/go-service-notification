package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
)

type Client interface {
	CreateClient(ctx context.Context, client *entity.Client) error
	UpdateClient(ctx context.Context, client *entity.Client) error
	DeleteClient(ctx context.Context, id string) error
}

type Notification interface {
	CreateNotification(ctx context.Context, notification *entity.Notification) error
	UpdateNotification(ctx context.Context, request *dto.UpdateNotificationRequest) error
	DeleteNotification(ctx context.Context, request *dto.DeleteNotificationRequest) error
	GetByCreateAt(ctx context.Context, request *dto.GetNotificationRequest) ([]entity.Notification, error)
}
