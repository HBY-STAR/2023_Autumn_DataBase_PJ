package commodity

import "github.com/gofiber/fiber/v2"

func RegisterRoutesWithoutAuthorization(app fiber.Router) {
	app.Get("/commodities/all", GetAllCommodity)
}

func RegisterRoutes(app fiber.Router) {
	app.Post("/commodities", CreateCommodity)
	app.Delete("/commodities/:id", DeleteCommodity)
	app.Put("/commodities", UpdateCommodity)
}
