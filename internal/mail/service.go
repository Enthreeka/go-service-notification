package mail

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/entity"
)

type MailService interface {
	GetTime(ctx context.Context) ([]entity.ClientsMessage, error)
	CreateMessageInfo(ctx context.Context, clientMessage *entity.ClientsMessage) error
}
