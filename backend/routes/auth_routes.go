package routes

import (
	"log"

	"github.com/fouradithep/todolist/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SetupAuthRoutes(app *fiber.App){
	if err := godotenv.Load(); err != nil {
		log.Fatal("load.env error")
	}
	app.Use("/tasks", handlers.AuthRequired)

}