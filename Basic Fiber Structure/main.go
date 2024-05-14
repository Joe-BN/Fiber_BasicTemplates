package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.dev/Joe-BN/Inter_Farm/app/routes"
)

func main() {
	//Loading Templates
	engine := html.New("./views", ".html")

	//Creating fiber app
    app := fiber.New(fiber.Config{
		Views: engine,
	})

	//configuring app
	app.Static("/", "./public")

	//Adding routes
	routes.SetupRoutes(app)

	// Start the App
    app.Listen(":3000")
}