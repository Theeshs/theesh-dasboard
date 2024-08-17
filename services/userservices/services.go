package userservices

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"theedashboard/ent"
)

func ListUserService(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	serviceID := c.Params("id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)
	if serviceID != "" {
		eduID, _ := strconv.ParseUint(serviceID, 10, 64)
		experience, err := FetchUserService(client, uint(uint64Value), uint(eduID))
		if err != nil {
			log.Fatalf("Error fetching user services: %s", err)
		}
		return c.JSON(experience)
	}
	services, err := FetchUsersServices(client, uint(uint64Value))

	if err != nil {
		log.Fatalf("Error fetching user services: %s", err)
	}

	return c.JSON(services)
}

func CreateUserService(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)
	service := new(UserService)

	if err := c.BodyParser(service); err != nil {
		println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdService, err := SaveUserService(client, uint(uint64Value), *service)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(createdService)
}

func UpdateUserService(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	serviceID := c.Params("id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)
	service := new(UserService)
	if serviceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to find service")
	}
	if err := c.BodyParser(service); err != nil {
		println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	servID, _ := strconv.ParseUint(serviceID, 10, 64)
	updatedService, err := EditUserService(client, uint(uint64Value), uint(servID), *service)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(updatedService)
}
