package priceChange

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Post("/price/history", GetPriceChange)
	app.Put("/price/history", UpdatePriceChange)
	//app.Post("/priceChange/search", SearchPriceChange)
}
