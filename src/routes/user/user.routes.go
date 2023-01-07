package user_routes

import (
	"github.com/gofiber/fiber/v2"
	user_controller "github.com/ssssshel/prueba-go-fiber/src/controllers/user"
)

func Routes(parent fiber.Router) fiber.Router {

	routesGroup := parent.Group("/users")

	routesGroup.Get("/get-users", user_controller.HandleGetUser)
	routesGroup.Post("/create-user", user_controller.HandleCreateUser)

	return routesGroup

}
