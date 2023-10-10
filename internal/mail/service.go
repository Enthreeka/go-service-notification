package mail

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"time"
)

type MailService interface {
	GetMail(ctx context.Context, t time.Time) ([]entity.ClientsMessage, error)
	CreateMessageInfo(ctx context.Context, clientMessage *entity.ClientsMessage) error
	GetSignal(ctx context.Context) ([]time.Time, error)
}
