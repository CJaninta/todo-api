package services

import "time"

type TodoResponse struct {
	ID        uint
	Title     string
	Detail    string
	Step      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoStepRequest struct {
	ID   int    `json:"id" validate:"required"`
	Step string `json:"step" validate:"required"`
}

type NewTodoRequest struct {
	Title     string `json:"title" validate:"required"`
	Detail    string `json:"detail" validate:"required"`
	Step      string `json:"step" validate:"required"`
	UserID    uint   `json:"user_id" validate:"required"`
	ProjectID uint   `json:"project_id" validate:"required"`
}

type TodoService interface {
	GetAllTodo(int) ([]TodoResponse, error)
	GetOneTodo(int) (*TodoResponse, error)
	UpdateTodoStep(TodoStepRequest) (*TodoResponse, error)
	CreateTodo(NewTodoRequest) (*TodoResponse, error)
}
