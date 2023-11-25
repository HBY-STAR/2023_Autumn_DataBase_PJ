package user

import (
	"github.com/gofiber/fiber/v2"
	. "src/models"
)

func GetCurrentUser(c *fiber.Ctx) error {
	user, err := GetUser(c)
	if err != nil {
		return err
	}
	return c.JSON(&user)
}

//func GetUserByID(c *fiber.Ctx) error {
//	userID, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	user, err := GetUser(c)
//	if err != nil {
//		return err
//	}
//
//	if !user.IsAdmin || user.ID == userID {
//		return common.Forbidden()
//	}
//
//	var getUser User
//	err = getUser.LoadUserByID(userID)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(&getUser)
//}
