package platform

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/platform/all", GetAllPlatform)
	app.Post("/platform", AddPlatform)
	app.Delete("/platform/:id", DeletePlatform)
	app.Put("/platform", UpdatePlatform)
}
