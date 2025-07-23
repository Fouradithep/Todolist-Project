package routes

import (
	"github.com/fouradithep/todolist/db"
	"github.com/fouradithep/todolist/handlers"
	"github.com/fouradithep/todolist/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SetupUserRoutes(app *fiber.App) {

	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(models.User)

		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err := handlers.CreatedUser(db.DB, user)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{"message": "Register Successful"})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(models.User)

		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		token, err := handlers.LoginUser(db.DB, user)

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
			Secure:   true,
			SameSite: fiber.CookieSameSiteNoneMode,
		})

		return c.JSON(fiber.Map{
			"message": "Login Successful",
		})

	})

	app.Post("/logout", func(c *fiber.Ctx) error {
    c.Cookie(&fiber.Cookie{
        Name:     "jwt",  // ชื่อ cookie ที่เคยเก็บ JWT token
        Value:    "",	// กำหนดให้เป็นค่าว่าง = ลบค่าเดิมออก
        Expires:  time.Now().Add(-time.Hour), // ตั้งเวลาให้หมดอายุไปแล้ว = ลบ cookie นี้
        HTTPOnly: true,
    })
    return c.JSON(fiber.Map{"message": "Logged out"})
})

}
