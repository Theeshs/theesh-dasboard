package experiences

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"theedashboard/ent"
)

func GetUserExperiences(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	experienceID := c.Params("id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)

	if experienceID != "" {
		expID, _ := strconv.ParseUint(experienceID, 10, 64)
		experience, err := FetchUserExperience(client, uint(uint64Value), uint(expID))
		if err != nil {
			log.Fatalf("Error fetching experiences: %s", err)
		}
		return c.JSON(experience)
	}

	experiences, err := FetchUserExperiences(client, uint(uint64Value))

	if err != nil {
		log.Fatalf("Error fetching experiences: %s", err)
	}

	return c.JSON(experiences)
}

func CreateUserExperience(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)

	// Create a new User struct
	experience := new(Experience)

	// Parse the JSON body into the User struct
	if err := c.BodyParser(experience); err != nil {
		println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdExp, err := CreateExperience(client, uint(uint64Value), *experience)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(createdExp)
}

func UpdateUserExperience(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	experienceID := c.Params("id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)
	experience := new(Experience)
	if experienceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to find experience")
	}

	// Parse the JSON body into the User struct
	if err := c.BodyParser(experience); err != nil {
		println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	expID, _ := strconv.ParseUint(experienceID, 10, 64)
	updatedExperience, err := UpdateExperience(client, uint(uint64Value), uint(expID), *experience)
	if err != nil {
		log.Fatalf("Error fetching experiences: %s", err)
	}

	return c.Status(fiber.StatusAccepted).JSON(updatedExperience)
}
