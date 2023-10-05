package repo

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/mail"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
)

type mailRepositoryPG struct {
	*postgres.Postgres
}

func NewMailRepositoryPG(postgres *postgres.Postgres) mail.MailStorage {
	return &mailRepositoryPG{
		postgres,
	}
}

func (m *mailRepositoryPG) GetCreateAt(ctx context.Context) (*entity.Notification, error) {
	query := `SELECT notification.id,notification.created_at,notification.message,notification.expires_at,
			   	   client.id,client.phone_number
			   FROM notification
			   		 JOIN client ON notification.id_client_properties = client.id_client_properties
			   WHERE notification.created_at = $1`

	// 1. В случае если pgx.ErrNoRows возвращать ошибку, далее ее выводить в лог в mail.Run()
	// 2. В случае если время совпало, то вернуться структуру(создать новую)
	// 3. Далее при рассылке сообщений передавать context.WithTimeout c expires_at

	return nil, nil
}
