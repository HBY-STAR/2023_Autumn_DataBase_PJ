package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	. "src/models"
)

func RegisterRoutes(app fiber.Router) {
	//app.Get("/users/me", GetCurrentUser)
	//app.Get("/users/:id", GetUserByID)
	//app.Put("/users/:id", ModifyUser)
	app.Post("/login", Login)
}

func Login(c *fiber.Ctx) error {
	// 解析请求体中的 JSON 数据
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&request); err != nil {
		return fiber.ErrBadRequest
	}

	//// 查询数据库中的用户
	//db, err := setupDatabase()
	//if err != nil {
	//	log.Fatal(err)
	//}
	var user User
	result := DB.Where("username = ?", request.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fiber.ErrUnauthorized
		}
		return fiber.ErrInternalServerError
	}

	// 验证密码
	if user.Password != request.Password {
		return fiber.ErrUnauthorized
	}

	// 登录成功
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Login successful",
	}
	return c.JSON(response)
}
