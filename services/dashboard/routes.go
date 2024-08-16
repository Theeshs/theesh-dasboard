package dashboard

import (
	"theedashboard/ent"
	"theedashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterDashboardRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return Home(c, client)
	})
}
