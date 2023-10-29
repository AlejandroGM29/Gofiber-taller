// En main.go

package main

import (
	"goFIber/config"

	"github.com/gofiber/fiber/v2"

	"goFIber/database"
)

func main() {
	app := fiber.New()

	database.Init()
	config.SetupRoutes(app)
	app.Listen("localhost:3000")

	defer database.Close()
}
