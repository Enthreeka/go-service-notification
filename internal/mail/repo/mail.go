package repo

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/mail"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"time"
)

type mailRepositoryPG struct {
	*postgres.Postgres
}

func NewMailRepositoryPG(postgres *postgres.Postgres) mail.MailStorage {
	return &mailRepositoryPG{
		postgres,
	}
}

func (m *mailRepositoryPG) CreateMessage(ctx context.Context, message *entity.Message) error {
	query := `INSERT INTO message (id,notification_id,client_id,created_at,status) VALUES ($1,$2,$3,$4,$5)`

	_, err := m.Pool.Exec(ctx, query,
		message.ID,
		message.NotificationID,
		message.ClientID,
		message.CreatedAt,
		message.Status)

	return err
}

func (m *mailRepositoryPG) GetMailing(ctx context.Context, t time.Time) ([]entity.ClientsMessage, error) {
	query := `SELECT notification.id, notification.created_at, notification.message, notification.expires_at,
					   client.id, client.phone_number
				FROM notification
				JOIN client ON notification.id_client_properties = client.id_client_properties
				WHERE notification.created_at = date_trunc('minute', $1::timestamp)`

	//currentTime := time.Now().Truncate(time.Minute)

	rows, err := m.Pool.Query(ctx, query, t)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	clientsMessage := make([]entity.ClientsMessage, 0)
	for rows.Next() {
		var clientMessage entity.ClientsMessage
		var createdAt time.Time
		var expiresAt time.Time

		err := rows.Scan(&clientMessage.NotificationID,
			&createdAt,
			&clientMessage.Message,
			&expiresAt,
			&clientMessage.ClientID,
			&clientMessage.PhoneNumber)
		if err != nil {
			return nil, err
		}

		clientMessage.ExpiresAt = expiresAt.String()
		clientMessage.CreatedAt = expiresAt.String()
		clientsMessage = append(clientsMessage, clientMessage)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return clientsMessage, nil
}

func (m *mailRepositoryPG) GetCreatedAt(ctx context.Context) ([]time.Time, error) {
	query := `SELECT created_at FROM signal`

	rows, err := m.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	times := make([]time.Time, 0)
	for rows.Next() {
		var t time.Time

		err := rows.Scan(&t)
		if err != nil {
			return nil, err
		}

		times = append(times, t)

	}

	if rows.Err() != nil {
		return nil, err
	}

	return times, nil
}
