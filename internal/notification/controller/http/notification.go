package http

import (
	"context"
	"errors"
	_ "github.com/Enthreeka/go-service-notification/docs"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type notificationHandler struct {
	notificationUsecase notification.NotificationService

	log *logger.Logger
}

func NewNotificationHandler(notificationUsecase notification.NotificationService, log *logger.Logger) *notificationHandler {
	return &notificationHandler{
		notificationUsecase: notificationUsecase,
		log:                 log,
	}
}

// CreateNotificationHandler godoc
// @Summary Create Notification
// @Tags notification
// @Description create notification
// @Accept json
// @Produce json
// @Param input body dto.CreateNotificationRequest true "Client new notification"
// @Success 201
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/notification/create [post]
func (n *notificationHandler) CreateNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &dto.CreateNotificationRequest{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = n.notificationUsecase.CreateNotification(context.Background(), notificationRequest)
	if err != nil {
		if errors.Is(err, apperror.ErrIncorrectTime) {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrIncorrectTime)
		} else if errors.Is(err, apperror.ErrClientAttribute) {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrClientAttribute)
		}

		n.log.Error("create notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

// UpdateNotificationHandler godoc
// @Summary Update Notification
// @Tags notification
// @Description update notification
// @Accept json
// @Produce json
// @Param input body dto.UpdateNotificationRequest true "Update already created notification"
// @Success 204
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/notification/update [post]
func (n *notificationHandler) UpdateNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &dto.UpdateNotificationRequest{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = n.notificationUsecase.UpdateNotification(context.Background(), notificationRequest)
	if err != nil {
		if err == apperror.ErrIncorrectTime {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrIncorrectTime)
		}
		n.log.Error("update notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteNotificationHandler godoc
// @Summary Delete Notification
// @Tags notification
// @Description delete notification
// @Accept json
// @Produce json
// @Param input body dto.TimeNotificationRequest true "Delete notification by his created time"
// @Success 204
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/notification/delete [delete]
func (n *notificationHandler) DeleteNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &dto.TimeNotificationRequest{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = n.notificationUsecase.DeleteNotification(context.Background(), notificationRequest)
	if err != nil {
		n.log.Error("delete notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetStatNotificationHandler godoc
// @Summary get Notification
// @Tags notification
// @Description get notification
// @Accept json
// @Produce json
// @Param input body dto.TimeNotificationRequest true "Get info by his created time about a specific notification "
// @Success 200 {object} []entity.Notification
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/notification/:time [Post]
func (n *notificationHandler) GetStatNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &dto.TimeNotificationRequest{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	stat, err := n.notificationUsecase.GetByCreateTime(context.Background(), notificationRequest)
	if err != nil {
		n.log.Error("get notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(stat)
}
