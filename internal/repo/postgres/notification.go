package postgres

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/repo"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"time"
)

type notificationRepositoryPG struct {
	*postgres.Postgres
}

func NewNotificationRepositoryPG(postgres *postgres.Postgres) repo.Notification {
	return &notificationRepositoryPG{
		postgres,
	}
}

func (n *notificationRepositoryPG) Create(ctx context.Context, notification *entity.Notification) error {
	query := `INSERT INTO notification
    (phone_number, operator_code, created_at, message, expires_at) 
			VALUES ($1,$2,$3,$4,$5)`

	expiresAtTime, err := time.Parse("15:04 02.01.2006", notification.ExpiresAt)
	if err != nil {
		return err
	}

	createdAtTime, err := time.Parse("15:04 02.01.2006", notification.CreateAt)
	if err != nil {
		return err
	}

	for _, property := range notification.ClientProperty {
		_, err = n.Pool.Exec(ctx, query,
			property.PhoneNumber,
			property.OperatorCode,
			createdAtTime,
			notification.Message,
			expiresAtTime,
		)
	}
	return err
}

func (n *notificationRepositoryPG) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM notification WHERE id = $1`

	_, err := n.Pool.Exec(ctx, query, id)
	return err
}
