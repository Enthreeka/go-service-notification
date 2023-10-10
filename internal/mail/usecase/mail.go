package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/mail"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
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

func (m *mailUsecase) GetMail(ctx context.Context, t time.Time) ([]entity.ClientsMessage, error) {
	if t.IsZero() {
		t = time.Now().Truncate(time.Minute)
	}

	clientsMessage, err := m.mailRepo.GetMailing(ctx, t)
	if err != nil {
		return nil, apperror.NewError("failed with getting method", err)
	}

	return clientsMessage, nil
}

func (m *mailUsecase) CreateMessageInfo(ctx context.Context, clientMessage *entity.ClientsMessage) error {
	message := &entity.Message{
		ID:             clientMessage.ID,
		NotificationID: clientMessage.NotificationID,
		ClientID:       clientMessage.ClientID,
		CreatedAt:      time.Now(),
	}
	if clientMessage.InTime == false {
		message.Status = "didn't have enough time"
	} else {
		message.Status = clientMessage.Status
	}

	err := m.mailRepo.CreateMessage(ctx, message)
	if err != nil {
		return apperror.NewError("failed to create message", err)
	}

	return nil
}

func (m *mailUsecase) GetSignal(ctx context.Context) ([]time.Time, error) {
	t, err := m.mailRepo.GetCreatedAt(ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}
