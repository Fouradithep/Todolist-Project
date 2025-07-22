package main

import (
	"fmt"
	"github.com/fouradithep/todolist/db"
	"github.com/fouradithep/todolist/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	db.Init()
	fmt.Println("Server started...")
	// รอเพิ่ม code run server หรือ handler ต่อไป

	app := fiber.New()

	app.Use(cors.New(cors.Config{
    AllowOrigins:     "https://todolist-frontend-zf8m.onrender.com/login",  // ระบุ frontend origin ให้ชัดเจน
    AllowCredentials: true,
    AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
}))

	routes.SetupAuthRoutes(app)

	routes.SetupTaskRoutes(app)

	routes.SetupUserRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback ถ้าไม่เจอ
	}
	app.Listen(":" + port)
	
}
