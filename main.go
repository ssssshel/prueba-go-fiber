package main

import (
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Response struct {
	Message string
	Data    any
}

type User struct {
	Firstname string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "Angel",
		Lastname:  "Arteaga",
	}

	res := Response{
		Message: "Successfull request",
		Data:    user,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func handleGetCpuQuantity(c *fiber.Ctx) error {

	cpusQnt := runtime.NumCPU()

	numCpusMessage := fmt.Sprintf("Number of CPUs: %d", cpusQnt)

	res := Response{
		Message: "Successfull request",
		Data:    numCpusMessage,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}

	err := c.BodyParser(&user)

	if err != nil {
		res := Response{
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := Response{
		Message: "Successfull request",
		Data:    user,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/user", handleUser)

	app.Get("/cpus", handleGetCpuQuantity)

	app.Post("/users", handleCreateUser)

	app.Listen(":3000")

}
