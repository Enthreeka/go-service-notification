package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/mail"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/google/uuid"
	"time"
)

type mailUsecase struct {
	mailRepo mail.MailStorage

	log *logger.Logger
}

func NewMailUsecase(mailRepo mail.MailStorage, log *logger.Logger) mail.MailService {
	return &mailUsecase{
		mailRepo: mailRepo,
		log:      log,
	}
}

func (m *mailUsecase) GetTime(ctx context.Context) ([]entity.ClientsMessage, error) {
	clientsMessage, err := m.mailRepo.GetMailing(ctx)
	if err != nil {
		return nil, apperror.NewError("failed with getting method", err)
	}

	return clientsMessage, nil
}

func (m *mailUsecase) CreateMessageInfo(ctx context.Context, clientMessage *entity.ClientsMessage) error {
	message := &entity.Message{
		ID:             uuid.New().String(),
		NotificationID: clientMessage.NotificationID,
		ClientID:       clientMessage.ClientID,
		CreatedAt:      time.Now(),
		Status:         clientMessage.Status,
	}

	err := m.mailRepo.CreateMessage(ctx, message)
	if err != nil {
		return apperror.NewError("failed to create message", err)
	}

	return nil
}
