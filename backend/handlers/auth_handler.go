package handlers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	jwtSecretKey := os.Getenv("jwtSecretKey") //in .env

	if jwtSecretKey == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "JWT secret key not set",
		})
	}

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)

	// แปลง user_id เป็น uint หรือ int64 ก่อนเก็บ (โดยทั่วไป token จะเก็บเป็น float64)
	userIDFloat, ok := claim["user_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user_id in token"})
	}

	userID := uint(userIDFloat)
	c.Locals("user_id", userID)

	log.Printf("Authenticated user ID: %d", userID)


	return c.Next()
}
