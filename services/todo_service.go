package services

import (
	"todo/repositories"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
)

type todoService struct {
	todoRepo    repositories.TodoRepository
	userRepo    repositories.UserRepository
	projectRepo repositories.ProjectRepository
}

func NewTodoService(
	todoRepo repositories.TodoRepository,
	userRepo repositories.UserRepository,
	projectRepo repositories.ProjectRepository,
) TodoService {
	return &todoService{
		todoRepo:    todoRepo,
		userRepo:    userRepo,
		projectRepo: projectRepo,
	}
}

func (s todoService) GetAllTodo(projectId int) ([]TodoResponse, error) {
	allTodo, err := s.todoRepo.GetByProjectId(projectId)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusNotFound,
			Message: "todo in this project id not found",
		}
	}

	resTodo := []TodoResponse{}
	for _, value := range allTodo {
		resTodo = append(resTodo, TodoResponse{
			ID:        value.ID,
			Title:     value.Title,
			Detail:    value.Detail,
			Step:      value.Step,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return resTodo, nil
}

func (s todoService) GetOneTodo(todoId int) (*TodoResponse, error) {
	return &TodoResponse{}, nil
}

func (s todoService) UpdateTodoStep(todoReq TodoStepRequest) (*TodoResponse, error) {
	err := s.todoRepo.UpdateStep(todoReq.ID, todoReq.Step)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusInternalServerError,
			Message: "failed updating",
		}
	}

	todo, err := s.todoRepo.GetById(todoReq.ID)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusNotFound,
			Message: "todo id not found",
		}
	}

	updatedTodo := TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Detail:    todo.Detail,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return &updatedTodo, nil
}

func (s todoService) CreateTodo(newTodo NewTodoRequest) (*TodoResponse, error) {
	addTodo := repositories.Todo{
		Title:     newTodo.Title,
		Detail:    newTodo.Detail,
		Step:      newTodo.Step,
		UserID:    newTodo.UserID,
		ProjectID: newTodo.ProjectID,
	}

	//TODO: should optimize use associate
	_, err := s.userRepo.GetById(newTodo.UserID)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusNotFound,
			Message: "user id not found",
		}
	}

	//TODO: should optimize use associate
	_, err = s.projectRepo.GetById(newTodo.ProjectID)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusNotFound,
			Message: "project id not found",
		}
	}

	todo, err := s.todoRepo.Create(addTodo)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusInternalServerError,
			Message: "failed creation",
		}
	}

	resTodo := TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Detail:    todo.Detail,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return &resTodo, nil
}
