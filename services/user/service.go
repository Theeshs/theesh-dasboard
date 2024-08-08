package user

import (
	"strconv"
	"theedashboard/ent"

	"github.com/gofiber/fiber/v2"
)

var jwtSecret = []byte("2131ouidjskbnfiu134kb..12m")

func GetUsers(c *fiber.Ctx, client *ent.Client) error {
	users, err := FetchUsers(client)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(users)
}

func GetUsersByID(c *fiber.Ctx, client *ent.Client) error {
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	user, err := FetchUserByID(client, uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}

// GetUserByEmail handles the GET /users/email/:email route
// @Summary Get a user by email
// @Description Get user by email
// @ID get-user-by-email
// @Accept  json
// @Produce  json
// @Param   email  path    string  true        "User Email"
// @Success 200 {object} ent.User
// @Failure 404 {object} fiber.Map
// @Router /users/email/{email} [get]
func GetUserByEmail(c *fiber.Ctx, client *ent.Client) error {
	email := c.Params("email")
	user, err := FetchUserByEmail(client, email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}

func LoginUser(c *fiber.Ctx, client *ent.Client) error {
	var userLogin UserLogin
	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	user, err := FetchUserByEmail(client, userLogin.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	// Compare the password
	accessToekn, refreshToken, err := GetTokens(*user, userLogin.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	return c.JSON(TokenResponse{
		AccessToken:  accessToekn,
		RefreshToken: refreshToken,
	})

}
