package device_routes

import (
	"github.com/gofiber/fiber/v2"
	device_info_controller "github.com/ssssshel/prueba-go-fiber/src/controllers/device"
)

func Routes(parent fiber.Router) fiber.Router {
	routesGroup := parent.Group("/device")

	routesGroup.Get("/get-cpus", device_info_controller.HandleGetCpuQuantity)

	return routesGroup
}
