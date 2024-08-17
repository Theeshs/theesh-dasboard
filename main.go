package main

import (
	"log"
	"theedashboard/ent"
	"theedashboard/services/educations"
	"theedashboard/services/experiences"
	"theedashboard/services/userservices"

	"theedashboard/services/dashboard"
	"theedashboard/services/user"

	"entgo.io/ent/dialect"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

// @title Fiber Swagger Example API
// @version 1.0
// @description This is Thee Dashboard APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Use a specific domain in production
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	client, err := ent.Open(dialect.Postgres, "postgres://postgres:mysecretpassword@127.0.0.1:5432/thee.me?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool to create all schema resources
	// if err := client.Schema.Create(context.Background()); err != nil {
	//     log.Fatalf("failed creating schema resources: %v", err)
	// }

	registerRoutes(app, client)
	log.Fatal(app.Listen(":4000"))

}

func registerRoutes(app *fiber.App, client *ent.Client) {
	dashboard.RegisterDashboardRoutes(app, client)
	user.RegisterUserRoutes(app, client)
	experiences.RegisterExperienceRoutes(app, client)
	educations.RegisterEducationRoutes(app, client)
	userservices.RegisterUserService(app, client)
}
