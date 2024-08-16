package experiences

import (
	"github.com/gofiber/fiber/v2"
	"theedashboard/ent"
	"theedashboard/middleware"
)

func RegisterExperienceRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/experiences", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return GetUserExperiences(c, client)
	})
	app.Get("/experiences/:id", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return GetUserExperiences(c, client)
	})
	app.Post("/experiences", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return CreateUserExperience(c, client)
	})
	app.Put("/experiences/:id", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return UpdateUserExperience(c, client)
	})
}
