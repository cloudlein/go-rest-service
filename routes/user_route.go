package routes

import (
	"github.com/cloudlein/go-rest-service/controller"
	"github.com/cloudlein/go-rest-service/repository"
	"github.com/cloudlein/go-rest-service/service"
	"github.com/gofiber/fiber/v2"
)

func RoutesInit(r *fiber.App) {

	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)

	r.Get("/api/v1/users", ctrl.GetAllUsers)
	r.Post("/api/v1/users", controller.CreateUsers)
	r.Get("/api/v1/users/:id", ctrl.GetUserById)
	r.Patch("/api/v1/users/:id", controller.UpdateUserById)
	r.Delete("/api/v1/users/:id", controller.DeleteUserById)

}
