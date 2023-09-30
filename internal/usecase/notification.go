package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/repo"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
)

type notificationUsecase struct {
	notificationRepoPG repo.Notification

	log *logger.Logger
}

func NewNotificationUsecase(notificationRepoPG repo.Notification, log *logger.Logger) Notification {
	return &notificationUsecase{
		notificationRepoPG: notificationRepoPG,
		log:                log,
	}
}

func (n *notificationUsecase) CreateNotification(ctx context.Context, notification *entity.Notification) error {
	if !entity.IsCorrectTime(notification.ExpiresAt) || !entity.IsCorrectTime(notification.CreateAt) {
		return apperror.ErrIncorrectTime
	}

	err := n.notificationRepoPG.Create(ctx, notification)
	if err != nil {
		return apperror.NewError("failed to create notification in postgres", err)
	}

	return nil
}

func (n *notificationUsecase) DeleteNotification(ctx context.Context, id string) error {
	err := n.notificationRepoPG.Delete(ctx, id)
	if err != nil {
		return apperror.NewError("failed to delete notification from postgres", err)
	}

	return nil
}
