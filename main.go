package main

import (
	"github.com/AbdulrahmanDaud10/notes-go-fiber-rest-api/database"
	"github.com/AbdulrahmanDaud10/notes-go-fiber-rest-api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func main() {
	
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	route.SetupRoutes(app)

	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
}
