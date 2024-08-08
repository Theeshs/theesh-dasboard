package user

import (
	"context"
	"fmt"
	"theedashboard/ent"
	"theedashboard/ent/user"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func FetchUsers(client *ent.Client) ([]*ent.User, error) {
	return client.User.Query().All(context.Background())
}

func FetchUserByID(client *ent.Client, id uint) (*ent.User, error) {
	return client.User.Get(context.Background(), id)
}

// FetchUserByEmail retrieves a user by email from the database
func FetchUserByEmail(client *ent.Client, email string) (*ent.User, error) {
	return client.User.Query().Where(user.Email(email)).Only(context.Background())
}

func genarateToken(userId uint, email string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    fmt.Sprintf("%d", userId),
		"exp":   time.Now().Add(duration).Unix(),
	})
	return token.SignedString(jwtSecret)
}

func GetTokens(user ent.User, password string) (string, string, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", err
	}
	accessToken, err := genarateToken(user.ID, user.Email, time.Hour*24*2)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := genarateToken(user.ID, user.Email, time.Hour*24*7)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
