package user

import (
	"theedashboard/ent"
	"theedashboard/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/users", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return GetUsers(c, client)
	})
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return GetUsersByID(c, client)
	})
	app.Get("/users/email/:email", func(c *fiber.Ctx) error {
		return GetUserByEmail(c, client)
	})

	app.Post("/users/login", func(c *fiber.Ctx) error {
		return LoginUser(c, client)
	})
}
