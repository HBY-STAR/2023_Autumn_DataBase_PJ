package apis

import (
	"github.com/gofiber/fiber/v2"
	"src/apis/user"
)

func registerRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api")
		//return c.SendString("Hello, World 👋!")
	})
}

// RegisterRoutes registers the necessary routes to the app
func RegisterRoutes(app *fiber.App) {
	registerRoutes(app)

	group := app.Group("/api")
	group.Use(MiddlewareGetUser)
	user.RegisterRoutes(group)
}

func MiddlewareGetUser(c *fiber.Ctx) error {
	//userObject, err := models.GetUser(c)
	//if err != nil {
	//	return err
	//}
	//c.Locals("user", userObject)
	//if config.Config.AdminOnly {
	//	if !userObject.IsAdmin {
	//		return common.Forbidden()
	//	}
	//}
	return c.Next()
}
