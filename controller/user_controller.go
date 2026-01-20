package controller

import (
	"strconv"

	"github.com/cloudlein/go-rest-service/database"
	"github.com/cloudlein/go-rest-service/models"
	"github.com/cloudlein/go-rest-service/service"
	"github.com/cloudlein/go-rest-service/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	users, total, err := c.service.GetAllUsers(page, limit)
	if err != nil {
		return ctx.Status(500).JSON(
			utils.ErrorResponse("failed to get all users", err),
		)
	}

	return ctx.JSON(
		utils.SuccessWithPagination(users, page, limit, total, "Successfully to get all users"),
	)
}

func CreateUsers(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newUser := models.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "status internal server error",
		})
	}

	newUser.Password = hashPassword

	database.DB.Debug().Create(&newUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Created new User",
	})
}

func GetUserById(c *fiber.Ctx) error {
	var user models.User

	result := database.DB.Debug().First(&user, c.Params("id"))

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

func UpdateUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is invalid",
		})
	}

	payload := make(map[string]interface{})

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if len(payload) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "payload is empty",
		})
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", id).Updates(payload).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update user",
	})
}

func DeleteUserById(c *fiber.Ctx) error {
	user := new(models.User)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Debug().Where("id = ?", id).Delete(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete user",
	})
}
