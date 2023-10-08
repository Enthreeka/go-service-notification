package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
)

type messageUsecase struct {
	messageRepoPG notification.MessageStorage

	log *logger.Logger
}

func NewMessageUsecase(messageRepoPG notification.MessageStorage, log *logger.Logger) notification.MessageService {
	return &messageUsecase{
		messageRepoPG: messageRepoPG,
		log:           log,
	}
}

func (m *messageUsecase) GetInfoNotification(ctx context.Context, id string) (map[string][]entity.MessageInfo, error) {
	message, err := m.messageRepoPG.GetAllByID(ctx, id)
	if err != nil {
		return nil, apperror.NewError("failed to get message by id", err)
	}

	return message, nil
}

func (m *messageUsecase) GetAllGroupByStatus(ctx context.Context) (map[string][]entity.MessageInfo, error) {
	message, err := m.messageRepoPG.GetAll(ctx)
	if err != nil {
		return nil, apperror.NewError("failed to get all message with group", err)
	}

	return message, nil
}
