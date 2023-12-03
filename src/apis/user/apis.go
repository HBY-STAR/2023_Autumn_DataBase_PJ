package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

//// GetCurrentUser @GetCurrentUser
//func GetCurrentUser(c *fiber.Ctx) error {
//	user, err := GetGeneralUser(c)
//	if err != nil {
//		return err
//	}
//	return c.JSON(&user)
//}

// GetUserInfo @GetUserInfo
// @Router /api/users/data [get]
// @Summary Get user info
// @Description Get user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Failure 403 {object} common.HttpError
// @Authorization Bearer {token}
func GetUserInfo(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "user" {
		return common.Forbidden()
	}

	getUser, err := GetUserByID(tmpUser.ID)
	if err != nil {
		return err
	}

	return c.JSON(&getUser)
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
// @Authorization Bearer {token}
func AddUser(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can add user")
	}
	var user User
	err = common.ValidateBody(c, &user)
	if err != nil {
		return err
	}
	return user.Create()
}

// UpdateUser @UpdateUser
// @Router /api/users [put]
// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param json body models.User true "json"
// @Success 200
// @Authorization Bearer {token}
func UpdateUser(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "user" {
		return common.Forbidden("Only admin or user can update user")
	}
	var user User
	err = common.ValidateBody(c, &user)
	if err != nil {
		return err
	}
	if tmpUser.UserType == "user" && tmpUser.ID != user.ID {
		return common.Forbidden("You can only update your own info")
	}
	return user.Update()
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
// @Authorization Bearer {token}
func DeleteUser(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can delete user")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	return DeleteUserByID(id)
}
