package message

import (
	"github.com/gofiber/fiber/v2"
	. "src/models"
)

// GetAllMessages @GetAllMessages
// @Router /api/message/all [get]
// @Summary Get all messages
// @Description Get all messages
// @Tags Message
// @Accept json
// @Produce json
// @Success 200 {object} Message
// @Authentication Bearer
func GetAllMessages(c *fiber.Ctx) error {
	var messages []Message
	return c.JSON(&messages)
}
