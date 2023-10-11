package postgres

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"strings"
	"time"
)

type notificationRepositoryPG struct {
	*postgres.Postgres
}

func NewNotificationRepositoryPG(postgres *postgres.Postgres) notification.NotificationStorage {
	return &notificationRepositoryPG{
		postgres,
	}
}

func (n *notificationRepositoryPG) CheckClientProperties(ctx context.Context, attributesMap map[string][]string) ([]string, error) {
	query := `SELECT id FROM client_properties WHERE tag = $1 AND operator_code = $2`

	var allID []string
	for key, value := range attributesMap {
		for _, attribute := range value {
			var id string

			err := n.Pool.QueryRow(ctx, query, key, attribute).Scan(&id)
			if err != nil {
				if err != pgx.ErrNoRows {
					return nil, err
				}
			} else {
				allID = append(allID, id)
			}

		}
	}

	return allID, nil
}

func (n *notificationRepositoryPG) Create(ctx context.Context, notification *entity.Notification) error {
	query := `INSERT INTO notification
    (id_client_properties, created_at, message, expires_at,with_signal) 
			VALUES ($1,$2,$3,$4,$5)`

	expiresAtTime, err := time.Parse("15:04 02.01.2006", notification.ExpiresAt)
	if err != nil {
		return err
	}

	createdAtTime, err := time.Parse("15:04 02.01.2006", notification.CreateAt)
	if err != nil {
		return err
	}

	for _, propertyId := range notification.ClientPropertyID {
		_, err = n.Pool.Exec(ctx, query,
			propertyId,
			createdAtTime,
			notification.Message,
			expiresAtTime,
			notification.Signal,
		)
	}

	return err
}

func (n *notificationRepositoryPG) Update(ctx context.Context, notification *entity.Notification) error {
	var counter int
	args := []interface{}{}
	var builder strings.Builder

	builder.WriteString("UPDATE notification SET")

	if notification.CreateAt != "" {
		counter++
		builder.WriteString(fmt.Sprintf(" created_at = $%d", counter))
		createdAtTime, err := time.Parse("15:04 02.01.2006", notification.CreateAt)
		if err != nil {
			return apperror.ErrIncorrectTime
		}

		createdAtFormatted := createdAtTime.Format("2006-01-02 15:04:05")
		args = append(args, createdAtFormatted)
	}

	if notification.Message != "" {
		counter++
		builder.WriteString(fmt.Sprintf(" ,message = $%d", counter))
		args = append(args, notification.Message)
	}

	if notification.ExpiresAt != "" {
		counter++
		builder.WriteString(fmt.Sprintf(" ,expires_at = $%d", counter))
		expiresAtTime, err := time.Parse("15:04 02.01.2006", notification.ExpiresAt)
		if err != nil {
			return apperror.ErrIncorrectTime
		}

		args = append(args, expiresAtTime)
	}

	counter++
	builder.WriteString(fmt.Sprintf(" WHERE created_at = $%d", counter))
	args = append(args, notification.CreateAt)

	_, err := n.Pool.Exec(ctx, builder.String(), args...)

	return err
}

func (n *notificationRepositoryPG) Delete(ctx context.Context, createdAt time.Time) error {
	query := `DELETE FROM notification WHERE created_at = $1`

	_, err := n.Pool.Exec(ctx, query, createdAt)
	return err
}

func (n *notificationRepositoryPG) GetByCreateAt(ctx context.Context, createdAt time.Time) ([]entity.Notification, error) {
	query := `SELECT notification.message,
				   notification.expires_at,
				   json_agg(json_build_object('tag', client_properties.tag, 'operator_code', client_properties.operator_code)) AS client_properties
					FROM
					notification
					JOIN client_properties ON notification.id_client_properties = client_properties.id
			WHERE
					created_at = $1
			GROUP BY
				 message, created_at, expires_at;`

	rows, err := n.Pool.Query(ctx, query, createdAt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notificationStat := make([]entity.Notification, 0)
	for rows.Next() {
		var notification entity.Notification
		var expiresAtTime time.Time

		err = rows.Scan(&notification.Message, &expiresAtTime, &notification.ClientProperty)
		if err != nil {
			return nil, err
		}

		notification.ExpiresAt = expiresAtTime.String()
		notification.CreateAt = createdAt.String()
		notificationStat = append(notificationStat, notification)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notificationStat, nil
}
