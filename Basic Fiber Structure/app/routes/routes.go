//Defining all my routes here

package routes

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.Render("homepage", fiber.Map{
			"Title": "Hello, World!",
		})
	})
}