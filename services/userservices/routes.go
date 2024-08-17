package userservices

import (
	"github.com/gofiber/fiber/v2"
	"theedashboard/ent"
	"theedashboard/middleware"
)

func RegisterUserService(app *fiber.App, client *ent.Client) {
	app.Get("/user-service", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return ListUserService(c, client)
	})
	app.Get("/user-service/:id", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return ListUserService(c, client)
	})
	app.Post("/user-service", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return CreateUserService(c, client)
	})
	app.Put("/user-service/:id", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return UpdateUserService(c, client)
	})
}
