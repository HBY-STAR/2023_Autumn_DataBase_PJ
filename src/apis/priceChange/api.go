package priceChange

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

// GetPriceChange @GetPriceChangeById
// @Router /api/price/history [post]
// @Summary Get priceChange by ID
// @Description Get priceChange by ID
// @Tags PriceChange
// @Accept json
// @Produce json
// @Param json body GetPriceChangeModel true "json"
// @Success 200 {object} PriceChange
// @Authentication Bearer
func GetPriceChange(c *fiber.Ctx) error {
	var getPriceChangeModel GetPriceChangeModel
	if err := c.BodyParser(&getPriceChangeModel); err != nil {
		// If there's an error in parsing, return an error response
		return common.BadRequest("Invalid request body")
	}
	priceChanges, err := GetPriceChangeById(getPriceChangeModel.CommodityItemID, getPriceChangeModel.TimeStart.Time, getPriceChangeModel.TimeEnd.Time)
	if err != nil {
		return err
	}
	return c.JSON(&priceChanges)
}

// UpdatePriceChange @UpdatePriceChange
// @Router /api/price/history [put]
// @Summary Update priceChange
// @Description Update priceChange
// @Tags PriceChange
// @Accept json
// @Produce json
// @Param json body UpdatePriceChangeModel true "json"
// @Success 200
// @Authentication Bearer
func UpdatePriceChange(c *fiber.Ctx) error {
	return nil
}
