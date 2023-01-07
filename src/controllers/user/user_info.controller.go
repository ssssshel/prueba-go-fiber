package user_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	global_interfaces "github.com/ssssshel/prueba-go-fiber/src/utils/interfaces"
)

func HandleGetUser(c *fiber.Ctx) error {

	user := global_interfaces.User{
		Id:        uuid.NewString(),
		Firstname: "Angel",
		Lastname:  "Arteaga",
	}

	res := global_interfaces.Response{
		Message: "Successfull request",
		Data:    user,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func HandleCreateUser(c *fiber.Ctx) error {

	user := global_interfaces.User{}

	err := c.BodyParser(&user)

	if err != nil {
		res := global_interfaces.Response{
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	user.Id = uuid.NewString()

	res := global_interfaces.Response{
		Message: "Successfull request",
		Data:    user,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
