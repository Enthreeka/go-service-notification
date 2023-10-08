package postgres

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"time"
)

type messageRepositoryPG struct {
	*postgres.Postgres
}

func NewMessageRepositoryPG(postgres *postgres.Postgres) notification.MessageStorage {
	return &messageRepositoryPG{
		postgres,
	}
}

func (m *messageRepositoryPG) GetAllByID(ctx context.Context, id string) (map[string][]entity.MessageInfo, error) {
	query := `SELECT message.created_at,message.status,client.phone_number,
				   client_properties.tag,client_properties.operator_code,notification.message
				FROM message
				JOIN client ON message.client_id = client.id
				JOIN notification ON message.notification_id = notification.id
				JOIN client_properties ON client.id_client_properties=client_properties.id
				WHERE message.id = $1;`

	rows, err := m.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	messages := make(map[string][]entity.MessageInfo, 0)
	for rows.Next() {
		var info entity.MessageInfo
		var createdAt time.Time

		err := rows.Scan(&createdAt,
			&info.Status,
			&info.PhoneNumber,
			&info.Tag,
			&info.OperatorCode,
			&info.Message,
		)
		if err != nil {
			return nil, err
		}

		info.CreatedAt = createdAt.String()

		if _, ok := messages[id]; !ok {
			messages[id] = append(messages[id], info)
		} else {
			messages[id] = append(messages[id], info)
		}
	}

	if rows.Err() != nil {
		return nil, err
	}

	return messages, nil
}

func (m *messageRepositoryPG) GetAll(ctx context.Context) (map[string][]entity.MessageInfo, error) {
	query := `SELECT notification.id, notification.message, message.status, COUNT(*)
				FROM message
				JOIN notification ON notification.id = message.notification_id
				GROUP BY
				notification.id, notification.message,message.status`

	rows, err := m.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	messages := make(map[string][]entity.MessageInfo, 0)
	for rows.Next() {
		var info entity.MessageInfo

		err := rows.Scan(&info.NotificationID,
			&info.Message,
			&info.Status,
			&info.Count,
		)
		if err != nil {
			return nil, err
		}

		if _, ok := messages[info.NotificationID]; !ok {
			messages[info.NotificationID] = append(messages[info.NotificationID], info)
		} else {
			messages[info.NotificationID] = append(messages[info.NotificationID], info)
		}
	}

	if rows.Err() != nil {
		return nil, err
	}

	return messages, nil
}
