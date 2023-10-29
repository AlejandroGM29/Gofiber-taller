// En config/config.go

package config

import (
	"goFIber/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/personas", controller.GetPersonas)
	app.Get("/personas/:id", controller.GetPersona)
	app.Post("/personas", controller.PostPersona)
}
