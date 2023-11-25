package favorite

import (
	"github.com/gofiber/fiber/v2"
	. "src/models"
)

// AddFavorite @AddFavorite
// @Router /api/favorites [post]
// @Summary AddFavorite
// @Description AddFavorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Param json body AddFavoriteModel true "json"
// @Success 200
// @Authentication Bearer
func AddFavorite(c *fiber.Ctx) error {
	return nil
}

// AddPriceLimit @AddPriceLimit
// @Router /api/price/limit [post]
// @Summary AddPriceLimit
// @Description AddPriceLimit
// @Tags Favorite
// @Accept json
// @Produce json
// @Param json body PriceLimitModel true "json"
// @Success 200
// @Authentication Bearer
func AddPriceLimit(c *fiber.Ctx) error {
	return nil
}

// GetAllFavorites @GetAllFavorites
// @Router /api/favorites/all [get]
// @Summary GetAllFavorites
// @Description GetAllFavorites
// @Tags Favorite
// @Accept json
// @Produce json
// @Success 200 {object} Favorite
// @Authentication Bearer
func GetAllFavorites(c *fiber.Ctx) error {
	var favorites []Favorite
	return c.JSON(&favorites)
}
