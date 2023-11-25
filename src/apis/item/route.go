package item

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/commodity/all", GetAllCommodity)
	app.Post("/search", SearchCommodity)
	app.Post("/commodity/item", AddCommodity)
	app.Put("/commodity/item", UpdateCommodity)
	app.Delete("/commodity/item/:id", DeleteCommodity)
}
