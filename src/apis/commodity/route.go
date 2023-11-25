package commodity

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/commodities/all", GetAllCommodity)
	app.Post("/commodities", AddCommodity)
	app.Delete("/commodities/:id", DeleteCommodity)
	app.Put("/commodities", UpdateCommodity)
}
