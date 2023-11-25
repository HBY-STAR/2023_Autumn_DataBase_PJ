package commodity

import (
	"github.com/gofiber/fiber/v2"
	. "src/models"
)

// GetAllCommodity @GetAllCommodity
// @Router /commodities/all [get]
// @Summary 获取所有商品
// @Description 获取所有商品
// @Tags 商品
// @Accept json
// @Produce json
// @Success 200 {object} Commodity
// @Authentication Bearer
func GetAllCommodity(c *fiber.Ctx) error {
	return nil
}

// AddCommodity @AddCommodity
// @Router /commodities [post]
// @Summary 添加商品
// @Description 添加商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param json body CreateCommodityRequest true "json"
// @Success 200
// @Authentication Bearer
func AddCommodity(c *fiber.Ctx) error {
	return nil
}

// DeleteCommodity @DeleteCommodity
// @Router /commodities [delete]
// @Summary 删除商品
// @Description 删除商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param id path string true "commodity id"
// @Success 200
// @Authentication Bearer
func DeleteCommodity(c *fiber.Ctx) error {
	return nil
}

// UpdateCommodity @UpdateCommodity
// @Router /commodities [put]
// @Summary 更新商品
// @Description 更新商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param json body Commodity true "json"
// @Success 200
// @Authentication Bearer
func UpdateCommodity(c *fiber.Ctx) error {
	var commodity Commodity
	err := c.BodyParser(&commodity)
	if err != nil {
		return err
	}
	return nil
}
