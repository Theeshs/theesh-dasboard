package educations

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"theedashboard/ent"
)

func GetUserExperiences(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	educationID := c.Params("id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)
	if educationID != "" {
		eduID, _ := strconv.ParseUint(educationID, 10, 64)
		experience, err := FetchUserEducation(client, uint(uint64Value), uint(eduID))
		if err != nil {
			log.Fatalf("Error fetching educations: %s", err)
		}
		return c.JSON(experience)
	}

	educations, err := FetchUserEducations(client, uint(uint64Value))
	if err != nil {
		log.Fatalf("Error fetching educations: %s", err)
	}

	return c.JSON(educations)
}
func CreateUserEducation(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)
	education := new(Education)

	if err := c.BodyParser(education); err != nil {
		println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	createdEducation, err := CreateEducation(client, uint(uint64Value), *education)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusCreated).JSON(createdEducation)
}

func UpdateUserEducation(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	educationId := c.Params("id")
	uint64Value, _ := strconv.ParseUint(userID.(string), 10, 64)
	education := new(Education)
	if educationId == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Unable to find experience")
	}
	if err := c.BodyParser(education); err != nil {
		println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	eduID, _ := strconv.ParseUint(educationId, 10, 64)
	updatedEducation, err := UpdateEducation(client, uint(uint64Value), uint(eduID), *education)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(updatedEducation)
}
