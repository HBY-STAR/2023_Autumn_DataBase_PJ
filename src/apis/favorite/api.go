package favorite

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
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
// @Authorization Bearer {token}
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
// @Authorization Bearer {token}
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
// @Authorization Bearer {token}
func GetAllFavorites(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "user" {
		return common.Forbidden("Only user can get favorites")
	}

	favorites, err := GetFavoritesByUserID(tmpUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(&favorites)
}
