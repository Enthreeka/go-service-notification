package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
	"github.com/Enthreeka/go-service-notification/internal/repo"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"time"
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

func (n *notificationUsecase) UpdateNotification(ctx context.Context, request *dto.UpdateNotificationRequest) error {

	t, err := time.Parse("", request.ExpiresAt)
	notification := &entity.Notification{
		Message:   request.Message,
		CreateAt:  request.CreateAt,
		ExpiresAt: t.String(),
	}

	err = n.notificationRepoPG.Update(ctx, notification)
	if err != nil {
		if err == apperror.ErrIncorrectTime {
			return apperror.ErrIncorrectTime
		}
		return apperror.NewError("failed to update notification in postgres", err)
	}

	return nil
}

func (n *notificationUsecase) DeleteNotification(ctx context.Context, request *dto.DeleteNotificationRequest) error {
	expiresAtTime, err := time.Parse("15:04 02.01.2006", request.CreateAt)
	if err != nil {
		return err
	}

	err = n.notificationRepoPG.Delete(ctx, expiresAtTime)
	if err != nil {
		return apperror.NewError("failed to delete notification from postgres", err)
	}

	return nil
}

func (n *notificationUsecase) GetByCreateAt(ctx context.Context, request *dto.GetNotificationRequest) ([]entity.Notification, error) {
	expiresAtTime, err := time.Parse("15:04 02.01.2006", request.CreateAt)
	if err != nil {
		return nil, err
	}

	notification, err := n.notificationRepoPG.GetByCreateAt(ctx, expiresAtTime)
	if err != nil {
		return nil, apperror.NewError("failed to get filtered notification from postgres", err)
	}

	return notification, nil
}
