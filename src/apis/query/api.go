package query

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

func GetFavoriteStatistics(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser == nil || tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can get favorite statistics")
	}
	var favoriteStatisticsRequest FavoriteStatisticsRequest
	if err := c.BodyParser(&favoriteStatisticsRequest); err != nil {
		return common.BadRequest("Invalid request body")
	}
	if favoriteStatisticsRequest.All {
		res, err := GetFavoriteStatisticsAll()
		if err != nil {
			return err
		}
		return c.JSON(res)
	}

	if favoriteStatisticsRequest.AgeStart > favoriteStatisticsRequest.AgeEnd || favoriteStatisticsRequest.AgeStart < 0 {
		return common.BadRequest("Invalid time range")
	}
	res, err := GetFavoriteStatisticsSome(favoriteStatisticsRequest.Gender, favoriteStatisticsRequest.AgeStart, favoriteStatisticsRequest.AgeEnd)
	if err != nil {
		return err
	}
	return c.JSON(res)
}
