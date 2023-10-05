package repo

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

type Client interface {
	Create(ctx context.Context, client *entity.Client) error
	Update(ctx context.Context, client *entity.Client) error
	Delete(ctx context.Context, id string) error
}

type Notification interface {
	CheckClientProperties(ctx context.Context, attributesMap map[string][]entity.Attribute) error
	Create(ctx context.Context, notification *entity.Notification) error
	Update(ctx context.Context, notification *entity.Notification) error
	Delete(ctx context.Context, createdAt time.Time) error
	GetByCreateAt(ctx context.Context, createdAt time.Time) ([]entity.Notification, error)
}

type Transaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
