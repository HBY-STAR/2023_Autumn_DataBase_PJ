package priceChange

import (
	"github.com/gofiber/fiber/v2"
	. "src/models"
)

// GetPriceChangeById @GetPriceChangeById
// @Router /api/price/history [post]
// @Summary Get priceChange by ID
// @Description Get priceChange by ID
// @Tags priceChange
// @Accept json
// @Produce json
// @Param json body GetPriceChangeModel true "json"
// @Success 200 {object} PriceChange
// @Authentication Bearer
func GetPriceChangeById(c *fiber.Ctx) error {
	var priceChanges []PriceChange
	return c.JSON(&priceChanges)
}

// UpdatePriceChange @UpdatePriceChange
// @Router /api/price/history [put]
// @Summary Update priceChange
// @Description Update priceChange
// @Tags priceChange
// @Accept json
// @Produce json
// @Param json body UpdatePriceChangeModel true "json"
// @Success 200
// @Authentication Bearer
func UpdatePriceChange(c *fiber.Ctx) error {
	return nil
}
