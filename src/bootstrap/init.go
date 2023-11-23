package bootstrap

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"src/apis"
	"src/models"
)

func Init() (*fiber.App, error) {
	err := models.InitDB()
	if err != nil {
		return nil, err
	}
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		//DisableStartupMessage: true,
	})
	apis.RegisterRoutes(app)
	return app, nil
}
