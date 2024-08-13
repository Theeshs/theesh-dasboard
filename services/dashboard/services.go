package dashboard

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"theedashboard/ent"
	"theedashboard/services/user"
)

func Home(c *fiber.Ctx, client *ent.Client) error {
	userID := c.Locals("user_id")
	println("User ID: ", userID)
	uint64Value, err := strconv.ParseUint(userID.(string), 10, 64)
	if err != nil {
		log.Fatalf("Error converting string to uint64: %v", err)
	}
	requestUser, err := user.FetchUserByID(client, uint(uint64Value))
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	return c.JSON(fiber.Map{
		"firstName": requestUser.FirstName,
		"username":  requestUser.Username,
		"message":   fmt.Sprintf("Welcome to %s Dashboard", requestUser.FirstName),
	})
}
