package message

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

// GetAllMessages @GetAllMessages
// @Router /api/messages [get]
// @Summary Get all messages
// @Description Get all messages
// @Tags Message
// @Accept json
// @Produce json
// @Success 200 {array} Message
// @Authentication Bearer
func GetAllMessages(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "user" {
		return common.Forbidden("Only user can get messages")
	}

	messages, err := GetMessagesByUserID(tmpUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(&messages)
}
