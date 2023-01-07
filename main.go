package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	device_routes "github.com/ssssshel/prueba-go-fiber/src/routes/device"
	user_routes "github.com/ssssshel/prueba-go-fiber/src/routes/user"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	// agrupacion de rutas de V1
	v1 := app.Group("/v1")

	user_routes.Routes(v1)
	device_routes.Routes(v1)

	app.Listen(":3000")

}
