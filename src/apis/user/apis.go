package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

// GetCurrentUser @GetCurrentUser
func GetCurrentUser(c *fiber.Ctx) error {
	user, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	return c.JSON(&user)
}

// GetUserInfo @GetUserInfo
// @Router /api/users/data [get]
// @Summary Get user info
// @Description Get user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Authentication Bearer
// @Failure 403 {object} common.HttpError
func GetUserInfo(c *fiber.Ctx) error {
	user, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	UserType := user.Type
	userID := user.ID

	if UserType != "user" {
		return common.Forbidden()
	}

	var getUser User
	err = getUser.LoadUserByID(userID)
	if err != nil {
		return err
	}

	return c.JSON(&user)
}

// AddUser @AddUser
// @Router /api/users [post]
// @Summary Add user
// @Description Add user
// @Tags User
// @Accept json
// @Produce json
// @Param json body CreateUserRequest true "json"
// @Success 200
// @Authentication Bearer
func AddUser(c *fiber.Ctx) error {
	return nil
}

// UpdateUser @UpdateUser
// @Router /api/users [put]
// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param json body UpdateUserRequest true "json"
// @Success 200
// @Authentication Bearer
func UpdateUser(c *fiber.Ctx) error {
	return nil
}

// DeleteUser @DeleteUser
// @Router /api/users [delete]
// @Summary Delete user
// @Description Delete user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200
// @Authentication Bearer
func DeleteUser(c *fiber.Ctx) error {
	return nil
}
