package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-vrf/src/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/generate-etcd-key", controller.GenerateEtcdKey)
	app.Post("/create-organization-vrf", controller.CreateOrganizationVRF)
}
