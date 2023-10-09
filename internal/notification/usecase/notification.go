package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"time"
)

type notificationUsecase struct {
	notificationRepoPG notification.NotificationStorage

	log *logger.Logger
}

func NewNotificationUsecase(notificationRepoPG notification.NotificationStorage, log *logger.Logger) notification.NotificationService {
	return &notificationUsecase{
		notificationRepoPG: notificationRepoPG,
		log:                log,
	}
}

func (n *notificationUsecase) CreateNotification(ctx context.Context, request *dto.CreateNotificationRequest) error {
	if !entity.IsCorrectTime(request.ExpiresAt) || !entity.IsCorrectTime(request.CreateAt) {
		return apperror.ErrIncorrectTime
	}

	if !entity.IsCorrectTime(request.ExpiresAt) || entity.IsCorrectTime(request.CreateAt) {
		return apperror.ErrIncorrectTime
	}

	notification := &entity.Notification{
		Message:   request.Message,
		CreateAt:  request.CreateAt,
		ExpiresAt: request.ExpiresAt,
	}

	id, err := n.existClientProperties(ctx, request)
	if err != nil {
		return apperror.NewError("failed check attribute properties", err)
	}

	if len(id) == 0 {
		return apperror.ErrClientAttribute
	}

	notification.ClientPropertyID = id

	err = n.notificationRepoPG.Create(ctx, notification)
	if err != nil {
		return apperror.NewError("failed to create notification in postgres", err)
	}

	return nil
}

func (n *notificationUsecase) existClientProperties(ctx context.Context, request *dto.CreateNotificationRequest) ([]string, error) {
	attributesMap := make(map[string][]string, len(request.Tags))

	for _, t := range request.Tags {
		for _, operator := range request.OperatorCodes {
			operatorCode := operator.OperatorCode

			if _, ok := attributesMap[t.Tag]; ok {
				attributesMap[t.Tag] = append(attributesMap[t.Tag], operatorCode)
			} else {
				attributesMap[t.Tag] = append(attributesMap[t.Tag], operatorCode)
			}
		}
	}

	id, err := n.notificationRepoPG.CheckClientProperties(ctx, attributesMap)
	if err != nil {
		return nil, err
	}

	return id, nil
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
