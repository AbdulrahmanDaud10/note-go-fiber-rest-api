package route

import (
	"github.com/AbdulrahmanDaud10/notes-go-fiber-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/user")

	// routes
	v1.Get("/", handlers.GetAllUsers)
	v1.Get("/:id", handlers.GetSingleUser)
	v1.Post("/", handlers.CreateUser)
	v1.Put("/:id", handlers.UpdateUser)
	v1.Delete("/:id", handlers.DeleteUserByID)
}
