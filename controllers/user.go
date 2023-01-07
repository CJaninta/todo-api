package controllers

import (
	"todo/services"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userSrv services.UserService
}

func NewUserController(userSrv services.UserService) userController {
	return userController{userSrv: userSrv}
}

func (uc userController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.userSrv.GetUsers()
	if err != nil {
		errs, ok := err.(utils.HandlerErr)
		if ok {
			return c.Status(errs.Code).JSON(utils.HandlerErr{
				Code:    errs.Code,
				Message: errs.Message,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (uc userController) GetOneUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "param bad request",
		})
	}

	user, err := uc.userSrv.GetOneUser(uint(id))
	if err != nil {
		errs, ok := err.(utils.HandlerErr)
		if ok {
			return c.Status(errs.Code).JSON(utils.HandlerErr{
				Code:    errs.Code,
				Message: errs.Message,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (uc userController) CreateUser(c *fiber.Ctx) error {
	user := services.NewUserRequest{}
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "body bad request",
		})
	}

	errors := utils.CheckRequest(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "request body required",
			Detail:  errors,
		})
	}

	newUser, err := uc.userSrv.CreateUser(user)
	if err != nil {
		errs, ok := err.(utils.HandlerErr)
		if ok {
			return c.Status(errs.Code).JSON(utils.HandlerErr{
				Code:    errs.Code,
				Message: errs.Message,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}
