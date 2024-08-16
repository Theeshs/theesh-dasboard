package educations

import (
	"github.com/gofiber/fiber/v2"
	"theedashboard/ent"
	"theedashboard/middleware"
)

func RegisterEducationRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/educations", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return GetUserExperiences(c, client)
	})
	app.Get("/educations/:id", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return GetUserExperiences(c, client)
	})
	app.Post("/educations/", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return CreateUserEducation(c, client)
	})
	app.Post("/educations/:id", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return UpdateUserEducation(c, client)
	})
}
