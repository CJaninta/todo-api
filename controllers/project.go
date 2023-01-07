package controllers

import (
	"fmt"
	"strconv"
	"todo/services"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
)

type projectController struct {
	projectSrv services.ProjectService
}

func NewProjectController(projectSrv services.ProjectService) projectController {
	return projectController{projectSrv: projectSrv}
}

func (pc projectController) GetProjects(c *fiber.Ctx) error {
	userId := c.Query("userId")
	userIdParsed, err := strconv.Atoi(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusNotFound,
			Message: "param bad request",
		})
	}

	projects, err := pc.projectSrv.GetProjects(userIdParsed)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.HandlerErr{
			Code:    fiber.StatusNotFound,
			Message: "projects not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(projects)
}

func (pc projectController) CreateProject(c *fiber.Ctx) error {
	project := services.NewProjectRequest{}
	err := c.BodyParser(&project)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "body bad request",
		})
	}

	errors := utils.CheckRequest(project)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "request body required",
			Detail:  errors,
		})
	}

	newProject, err := pc.projectSrv.CreateProject(project)
	if err != nil {
		errs, ok := err.(utils.HandlerErr)
		if ok {
			return c.Status(errs.Code).JSON(utils.HandlerErr{
				Code:    errs.Code,
				Message: errs.Message,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(newProject)
}

func (pc projectController) AddUserInProject(c *fiber.Ctx) error {

	userProject := services.UserProjectRequest{}
	err := c.BodyParser(&userProject)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "body bad request",
		})
	}

	errors := utils.CheckRequest(userProject)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "request body required",
			Detail:  errors,
		})
	}

	err = pc.projectSrv.AddUserInProject(userProject.UserId, userProject.ProjectId)
	if err != nil {
		errs, ok := err.(utils.HandlerErr)
		if ok {
			return c.Status(errs.Code).JSON(utils.HandlerErr{
				Code:    errs.Code,
				Message: errs.Message,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": fmt.Sprintf("add user id = %v in project id = %v success", userProject.UserId, userProject.ProjectId),
	})
}
