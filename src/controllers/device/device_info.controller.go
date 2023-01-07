package device_info_controller

import (
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
	global_interfaces "github.com/ssssshel/prueba-go-fiber/src/utils/interfaces"
)

func HandleGetCpuQuantity(c *fiber.Ctx) error {

	cpusQnt := runtime.NumCPU()

	numCpusMessage := fmt.Sprintf("Number of CPUs: %d", cpusQnt)

	res := global_interfaces.Response{
		Message: "Successfull request",
		Data:    numCpusMessage,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
