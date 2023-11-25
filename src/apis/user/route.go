package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/users/me", GetCurrentUser)
	app.Get("/users/data", GetUserInfo)
	app.Post("/users", AddUser)
	app.Delete("/users/:id", DeleteUser)
	app.Put("/users", UpdateUser)
}
