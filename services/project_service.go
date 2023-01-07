package services

import (
	"todo/repositories"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
)

type projectService struct {
	userRepo        repositories.UserRepository
	projectRepo     repositories.ProjectRepository
	userProjectRepo repositories.UserProjectRepository
}

func NewProjectService(
	userRepo repositories.UserRepository,
	projectRepo repositories.ProjectRepository,
	userProjectRepo repositories.UserProjectRepository,
) ProjectService {
	return &projectService{
		userRepo:        userRepo,
		projectRepo:     projectRepo,
		userProjectRepo: userProjectRepo,
	}
}

func (s projectService) GetProjects(userId int) ([]ProjectResponse, error) {
	projects, err := s.projectRepo.GetProjectByUser(userId)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusInternalServerError,
			Message: "internal error",
		}
	}

	resProjects := []ProjectResponse{}
	for _, value := range projects {
		resProjects = append(resProjects, ProjectResponse{
			ID:     value.ID,
			Name:   value.Name,
			Detail: value.Detail,
		})
	}

	return resProjects, nil
}

func (s projectService) GetUserInProject(projectId int) ([]UserResponse, error) { //TODO
	return []UserResponse{}, nil
}

func (s projectService) CreateProject(newProject NewProjectRequest) (*ProjectResponse, error) {
	addProject := repositories.Project{
		Name:   newProject.Name,
		Detail: newProject.Detail,
	}

	projectName, _ := s.projectRepo.GetByName(addProject.Name)
	if projectName != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusConflict,
			Message: "project already exists",
		}
	}

	project, err := s.projectRepo.Create(addProject)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusInternalServerError,
			Message: "failed creation",
		}
	}

	resProject := ProjectResponse{
		ID:     project.ID,
		Name:   project.Name,
		Detail: project.Detail,
	}

	return &resProject, nil
}

func (s projectService) AddUserInProject(userId int, projectId int) error {
	err := s.userProjectRepo.AddUserInProject(userId, projectId)
	if err != nil {
		return utils.HandlerErr{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}
