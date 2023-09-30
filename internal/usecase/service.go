package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
)

type Client interface {
	CreateClient(ctx context.Context, client *entity.Client) error
	UpdateClient(ctx context.Context, client *entity.Client) error
	DeleteClient(ctx context.Context, id string) error
}

type Notification interface {
	CreateNotification(ctx context.Context, notification *entity.Notification) error
	DeleteNotification(ctx context.Context, id string) error
}
