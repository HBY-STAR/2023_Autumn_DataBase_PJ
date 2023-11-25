package item

import (
	"github.com/gofiber/fiber/v2"
	. "src/models"
)

// GetAllCommodity @GetAllCommodity
// @Router /api/commodity/all [get]
// @Summary 获取所有商品
// @Description 获取所有商品
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {object} CommodityItem
// @Authentication Bearer
func GetAllCommodity(c *fiber.Ctx) error {
	var CommodityItems []CommodityItem
	err := DB.Find(&CommodityItems).Error
	if err != nil {
		return err
	}
	return c.JSON(&CommodityItems)
}

// SearchCommodity @SearchCommodity
// @Router /api/search [post]
// @Summary 搜索商品
// @Description 搜索商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body SearchQuery true "Range: name, category; Search: content to search"
// @Success 200 {object} CommodityItem
// @Authentication Bearer
func SearchCommodity(c *fiber.Ctx) error {
	var CommodityItems []CommodityItem
	//err := DB.Find(&CommodityItems).Error
	//if err != nil {
	//	return err
	//}
	return c.JSON(&CommodityItems)
}

// AddCommodity @AddCommodity
// @Router /api/commodity/item [post]
// @Summary 添加商品
// @Description 添加商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body CreateItemModel true "json"
// @Success 200
// @Authentication Bearer
func AddCommodity(c *fiber.Ctx) error {
	var commodityItem CommodityItem
	err := c.BodyParser(&commodityItem)
	if err != nil {
		return err
	}
	err = DB.Create(&commodityItem).Error
	if err != nil {
		return err
	}
	return c.JSON(&commodityItem)
}

// UpdateCommodity @UpdateCommodity
// @Router /api/commodity/item [put]
// @Summary 更新商品
// @Description 更新商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body UpdateItemModel true "json"
// @Success 200
// @Authentication Bearer
func UpdateCommodity(c *fiber.Ctx) error {
	var commodityItem CommodityItem
	err := c.BodyParser(&commodityItem)
	if err != nil {
		return err
	}
	err = DB.Save(&commodityItem).Error
	if err != nil {
		return err
	}
	return c.JSON(&commodityItem)
}

// DeleteCommodity @DeleteCommodity
// @Router /api/commodity/item [delete]
// @Summary 删除商品
// @Description 删除商品
// @Tags Item
// @Accept json
// @Produce json
// @Param id path string true "commodity id"
// @Success 200
// @Authentication Bearer
func DeleteCommodity(c *fiber.Ctx) error {
	return nil
}
