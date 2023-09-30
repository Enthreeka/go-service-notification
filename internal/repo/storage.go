package repo

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
)

type Client interface {
	Create(ctx context.Context, client *entity.Client) error
	Update(ctx context.Context, client *entity.Client) error
	Delete(ctx context.Context, id string) error
}

type Notification interface {
	Create(ctx context.Context, notification *entity.Notification) error
	Delete(ctx context.Context, id string) error
}
