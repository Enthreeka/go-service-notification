package mail

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
)

type MailStorage interface {
	GetCreateAt(ctx context.Context) (*entity.Notification, error)
}
