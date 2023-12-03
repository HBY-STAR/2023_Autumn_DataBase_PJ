package priceChange

import "github.com/gofiber/fiber/v2"

func RegisterRoutesWithoutAuthorization(app fiber.Router) {
	app.Post("/price/history", GetPriceChange)
}

func RegisterRoutes(app fiber.Router) {
	app.Put("/price/history", UpdatePriceChange)
	app.Post("/price/history/batch", AddBatchPriceChange)
	//app.Post("/priceChange/search", SearchPriceChange)
}
