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

func (n *notificationUsecase) CreateNotification(ctx context.Context, request *dto.CreateNotificationRequest) error {
	if !entity.IsCorrectTime(request.ExpiresAt) || !entity.IsCorrectTime(request.CreateAt) {
		return apperror.ErrIncorrectTime
	}

	notification := &entity.Notification{
		Message:   request.Message,
		CreateAt:  request.CreateAt,
		ExpiresAt: request.ExpiresAt,
	}

	multiple := make(map[string][]string, len(request.Tags)*len(request.OperatorCodes))

	for _, tag := range request.Tags {
		for _, code := range request.OperatorCodes {

			if storeTag, ok := multiple[tag.Tag]; ok {
				storeTag = append(storeTag, code.OperatorCode)
			} else {
				multiple[tag.Tag] = append(multiple[tag.Tag], code.OperatorCode)
			}

		}
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

func (n *notificationUsecase) DeleteNotification(ctx context.Context, request *dto.TimeNotificationRequest) error {
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

func (n *notificationUsecase) GetByCreateTime(ctx context.Context, request *dto.TimeNotificationRequest) ([]entity.Notification, error) {
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
