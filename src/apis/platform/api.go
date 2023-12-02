package platform

import (
	"github.com/gofiber/fiber/v2"
	. "src/models"
)

// GetAllPlatform @GetAllPlatform
// @Router /api/platform/all [get]
// @Summary Get all platform
// @Description Get all platform
// @Tags Platform
// @Accept json
// @Produce json
// @Success 200 {object} Platform
// @Authentication Bearer
func GetAllPlatform(c *fiber.Ctx) error {
	platforms, err := GetPlatforms()
	if err != nil {
		return err
	}
	return c.JSON(&platforms)
}

// AddPlatform @AddPlatform
// @Router /api/platform [post]
// @Summary Add platform
// @Description Add platform
// @Tags Platform
// @Accept json
// @Produce json
// @Param json body CreatePlatformRequest true "json"
// @Success 200
// @Authentication Bearer
func AddPlatform(c *fiber.Ctx) error {
	//var platform Platform
	//err := common.ValidateBody(c, &platform)
	//if err != nil {
	//	return err
	//}
	//
	//err = platform.Create()
	//if err != nil {
	//	return err
	//}
	//
	//return c.JSON(&platform)
	return nil
}

// UpdatePlatform @UpdatePlatform
// @Router /api/platform [put]
// @Summary Update platform
// @Description Update platform
// @Tags Platform
// @Accept json
// @Produce json
// @Param json body UpdatePlatformRequest true "json"
// @Success 200
// @Authentication Bearer
func UpdatePlatform(c *fiber.Ctx) error {
	return nil
}

// DeletePlatform @DeletePlatform
// @Router /api/platform [delete]
// @Summary Delete platform
// @Description Delete platform
// @Tags Platform
// @Accept json
// @Produce json
// @Param id path int true "platform id"
// @Success 200
// @Authentication Bearer
func DeletePlatform(c *fiber.Ctx) error {
	return nil
}
