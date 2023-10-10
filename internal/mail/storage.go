package mail

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"time"
)

type MailStorage interface {
	GetMailing(ctx context.Context, currentTime time.Time) ([]entity.ClientsMessage, error)
	CreateMessage(ctx context.Context, message *entity.Message) error
	GetCreatedAt(ctx context.Context) ([]time.Time, error)
}
