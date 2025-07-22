package routes

import (
	

	"github.com/fouradithep/todolist/handlers"

	"github.com/gofiber/fiber/v2"
	
)

func SetupAuthRoutes(app *fiber.App){
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("load.env error")
	// }
	app.Use("/tasks", handlers.AuthRequired)

}