package controllers

import (
	"todo/services"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
)

type todoController struct {
	todoSrv services.TodoService
}

func NewTodoController(todoSrv services.TodoService) todoController {
	return todoController{todoSrv: todoSrv}
}

func (tc todoController) CreateTodo(c *fiber.Ctx) error {
	todo := services.NewTodoRequest{}
	err := c.BodyParser(&todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "body bad request",
		})
	}

	errors := utils.CheckRequest(todo)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "request body required",
			Detail:  errors,
		})
	}

	newTodo, err := tc.todoSrv.CreateTodo(todo)
	if err != nil {
		errs, ok := err.(utils.HandlerErr)
		if ok {
			return c.Status(errs.Code).JSON(utils.HandlerErr{
				Code:    errs.Code,
				Message: errs.Message,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(newTodo)
}

func (tc todoController) UpdateTodoStep(c *fiber.Ctx) error {
	todo := services.TodoStepRequest{}
	err := c.BodyParser(&todo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "body bad request",
		})
	}

	errors := utils.CheckRequest(todo)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "request body required",
			Detail:  errors,
		})
	}

	updatedTodo, err := tc.todoSrv.UpdateTodoStep(todo)
	if err != nil {
		errs, ok := err.(utils.HandlerErr)
		if ok {
			return c.Status(errs.Code).JSON(utils.HandlerErr{
				Code:    errs.Code,
				Message: errs.Message,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(updatedTodo)
}
