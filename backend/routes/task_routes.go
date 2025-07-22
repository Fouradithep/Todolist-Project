package routes

import (
	"github.com/fouradithep/todolist/db"
	"github.com/fouradithep/todolist/handlers"
	"github.com/fouradithep/todolist/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func SetupTaskRoutes(app *fiber.App) {
	// Read All
	app.Get("/tasks", handlers.AuthRequired, func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		tasks := handlers.GetTasks(db.DB, userID)
		return c.JSON(tasks)
	})

	// Read
	app.Get("/task/:id", handlers.AuthRequired, func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}

		task, err := handlers.GetTask(db.DB, id, userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(task)
	})

	// Create
	app.Post("/task", handlers.AuthRequired, func(c *fiber.Ctx) error {
		task := new(models.Task)
		if err := c.BodyParser(task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse body"})
		}

		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		err := handlers.CreateTask(db.DB, task, userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{
			"message": "Create Task Successful",
			"id": task.ID,
			"title": task.Title,
			"status":task.Status,
		})
	})

	// Update
	app.Put("/task/:id", handlers.AuthRequired, func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}

		task := new(models.Task)
		if err := c.BodyParser(task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse body"})
		}

		task.ID = uint(id)
		err = handlers.UpdateTask(db.DB, task, userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"message": "Update Task Successful"})
	})

	// Delete
	app.Delete("/task/:id", handlers.AuthRequired, func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid task ID"})
		}

		err = handlers.DeletedTask(db.DB, id, userID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"message": "Delete Task Successful"})
	})
}
