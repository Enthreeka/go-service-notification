package http

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/usecase"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type notificationHandler struct {
	notificationUsecase usecase.Notification

	log *logger.Logger
}

func NewNotificationHandler(notificationUsecase usecase.Notification, log *logger.Logger) *notificationHandler {
	return &notificationHandler{
		notificationUsecase: notificationUsecase,
		log:                 log,
	}
}

func (n *notificationHandler) CreateNotificationHandler(c *fiber.Ctx) error {
	notification := &entity.Notification{}

	err := c.BodyParser(notification)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = n.notificationUsecase.CreateNotification(context.Background(), notification)
	if err != nil {
		if err == apperror.ErrIncorrectTime {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrIncorrectTime)
		}
		n.log.Error("create notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (n *notificationHandler) DeleteNotificationHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	err := n.notificationUsecase.DeleteNotification(context.Background(), id)
	if err != nil {
		n.log.Error("delete notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
