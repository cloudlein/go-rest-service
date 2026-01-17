package routes

import (
	"github.com/cloudlein/go-rest-service/controller"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(r *fiber.App) {

	r.Get("/api/v1/users", controller.GetAllUsers)
	r.Post("/api/v1/users", controller.CreateUsers)
	r.Get("/api/v1/users/:id", controller.GetUserById)
	r.Patch("/api/v1/users/:id", controller.UpdateUserById)
	r.Delete("/api/v1/users/:id", controller.DeleteUserById)

}
