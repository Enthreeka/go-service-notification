package http

import (
	"context"
	_ "github.com/Enthreeka/go-service-notification/docs"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type messageHandler struct {
	messageUsecase notification.MessageService

	log *logger.Logger
}

func NewMessageHandler(messageUsecase notification.MessageService, log *logger.Logger) *messageHandler {
	return &messageHandler{
		messageUsecase: messageUsecase,
		log:            log,
	}
}

// GetDetailInfoHandler godoc
// @Summary Get Messages
// @Tags message
// @Description get message by id
// @Accept json
// @Produce json
// @Param id path string true "ID of the message"
// @Success 200 {object} map[string][]entity.MessageInfo
// @Failure 400 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/message/info/{id} [get]
func (m *messageHandler) GetDetailInfoHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	m.log.Info("%s", id)

	message, err := m.messageUsecase.GetInfoNotification(context.Background(), id)
	if err != nil {
		m.log.Error("get info controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	if len(message) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(apperror.ErrNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(message)
}

// GetGroupByStatusHandler godoc
// @Summary Get Messages Group By
// @Tags message
// @Description get all messages
// @Produce json
// @Success 200 {object} map[string][]entity.MessageInfo
// @Failure 500 {object} apperror.AppError
// @Router /api/message/group [get]
func (m *messageHandler) GetGroupByStatusHandler(c *fiber.Ctx) error {
	message, err := m.messageUsecase.GetAllGroupByStatus(context.Background())
	if err != nil {
		m.log.Error("get info group by controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(message)
}
