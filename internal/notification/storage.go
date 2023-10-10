package notification

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"time"
)

type ClientStorage interface {
	Create(ctx context.Context, client *entity.Client) error
	Update(ctx context.Context, client *entity.Client) error
	Delete(ctx context.Context, id string) error
}

type NotificationStorage interface {
	CheckClientProperties(ctx context.Context, attributesMap map[string][]string) ([]string, error)
	Create(ctx context.Context, notification *entity.Notification) error
	Update(ctx context.Context, notification *entity.Notification) error
	Delete(ctx context.Context, createdAt time.Time) error
	GetByCreateAt(ctx context.Context, createdAt time.Time) ([]entity.Notification, error)
}

type MessageStorage interface {
	GetAllByID(ctx context.Context, id string) (map[string][]entity.MessageInfo, error)
	GetAll(ctx context.Context) (map[string][]entity.MessageInfo, error)
}

type Signal interface {
	Create(ctx context.Context, time string) error
}
