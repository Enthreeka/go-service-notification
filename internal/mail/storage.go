package mail

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
)

type MailStorage interface {
	GetMailing(ctx context.Context) ([]entity.ClientsMessage, error)
	CreateMessage(ctx context.Context, message *entity.Message) error
}
