package auth

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	//app.Get("/users/me", GetCurrentUser)
	//app.Get("/users/:id", GetUserByID)
	//app.Put("/users/:id", ModifyUser)
	app.Post("/login", Login)
}