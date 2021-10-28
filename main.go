package main

import (
	"test/waldhalf/gofiber-tpl/src/database"
	"test/waldhalf/gofiber-tpl/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

    app := fiber.New()
	
	// Enable cors
	app.Use(cors.New(cors.Config{AllowCredentials: true}))


	routes.Setup(app)
	
    app.Listen(":8000")
}