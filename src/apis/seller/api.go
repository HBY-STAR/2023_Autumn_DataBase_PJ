package seller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

//// GetCurSeller @GetCurSeller
//// @Router /api/sellers/me [get]
//func GetCurSeller(ctx *fiber.Ctx) error {
//	return nil
//}

// GetSeller @GetSeller
// @Router /api/sellers/data [get]
// @Summary Get seller info
// @Description Get seller info
// @Tags Seller
// @Accept json
// @Produce json
// @Success 200 {object} Seller
// @Authentication Bearer
// @Failure 403 {object} common.HttpError
func GetSeller(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "seller" {
		return common.Forbidden("You are not a seller.")
	}

	seller, err := GetSellerByID(tmpUser.ID)
	if err != nil {
		return err
	}

	return c.JSON(&seller)
}

// AddSeller @AddSeller
// @Router /api/sellers [post]
// @Summary Add seller
// @Description Add seller
// @Tags Seller
// @Accept json
// @Produce json
// @Param json body CreateSellerRequest true "json"
// @Success 200
// @Authentication Bearer
func AddSeller(c *fiber.Ctx) error {
	//var seller Seller
	//err := common.ValidateBody(c, &seller)
	//if err != nil {
	//	return err
	//}
	//
	//err = seller.Create()
	//if err != nil {
	//	return err
	//}
	//
	//return c.JSON(seller)
	return nil
}

// UpdateSeller @UpdateSeller
// @Router /api/sellers [put]
// @Summary Update seller
// @Description Update seller
// @Tags Seller
// @Accept json
// @Produce json
// @Param json body UpdateSellerRequest true "json"
// @Success 200
// @Authentication Bearer
func UpdateSeller(c *fiber.Ctx) error {
	return nil
}

// DeleteSeller @DeleteSeller
// @Router /api/sellers [delete]
// @Summary Delete seller
// @Description Delete seller
// @Tags Seller
// @Accept json
// @Produce json
// @Param id path int true "seller id"
// @Success 200
// @Authentication Bearer
func DeleteSeller(c *fiber.Ctx) error {
	return nil
}
