package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go-vrf/src/routes"
)

func main() {

	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen(":4000")
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
		return
	}
}
