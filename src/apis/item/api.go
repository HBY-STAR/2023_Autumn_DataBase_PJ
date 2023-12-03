package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
	"strings"
)

// GetAllCommodity @GetAllCommodity
// @Router /api/commodity/all [get]
// @Summary 获取所有商品
// @Description 获取所有商品
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {array} models.CommodityItem
// @Failure 403 {object} common.HttpError
func GetAllCommodity(c *fiber.Ctx) error {
	items, err := GetItems()
	if err != nil {
		return err
	}
	return c.JSON(&items)
}

// SearchCommodity @SearchCommodity
// @Router /api/search [post]
// @Summary 搜索商品
// @Description 搜索商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body SearchQuery true "Range: name, category; Search: content to search"
// @Success 200 {array} models.CommodityItem
func SearchCommodity(c *fiber.Ctx) error {
	var searchQuery SearchQuery
	err := c.BodyParser(&searchQuery)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	var items []CommodityItem
	switch strings.ToLower(searchQuery.Range) {
	case "name":
		if searchQuery.Accurate {
			items, err = GetItemsByName(searchQuery.Search)
		} else {
			items, err = GetItemsByNameFuzzy(searchQuery.Search)
		}
	case "category":
		if searchQuery.Accurate {
			items, err = GetItemsByCategory(searchQuery.Search)
		} else {
			items, err = GetItemsByCategoryFuzzy(searchQuery.Search)
		}
	default:
		return common.BadRequest("Invalid search range")
	}
	if err != nil {
		return err
	}
	return c.JSON(&items)
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
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Authorization Bearer {token}
func AddCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can add items")
	}
	var createItemModel CreateItemModel
	err = c.BodyParser(&createItemModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	// check if valid
	var commodityItem = CommodityItem{
		ItemName:    createItemModel.ItemName,
		Price:       createItemModel.Price,
		SellerID:    tmpUser.ID,
		PlatformID:  createItemModel.PlatformID,
		CommodityID: createItemModel.CommodityID,
	}

	return commodityItem.Create()
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
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Authorization Bearer {token}
func UpdateCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can update items")
	}
	var updateItemModel UpdateItemModel
	err = c.BodyParser(&updateItemModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	var commodityItem = CommodityItem{
		ID:       updateItemModel.CommodityItemID,
		ItemName: updateItemModel.ItemName,
		Price:    updateItemModel.Price,
	}
	return commodityItem.Update()
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
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Authorization Bearer {token}
func DeleteCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can delete items")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	return DeleteItemByID(id)
}
