package services

type ProjectResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type NewProjectRequest struct {
	Name   string `validate:"required"`
	Detail string `validate:"required"`
}

type UserProjectRequest struct {
	UserId    int `json:"user_id" validate:"required"`
	ProjectId int `json:"project_id" validate:"required"`
}

type ProjectService interface {
	GetProjects(int) ([]ProjectResponse, error)
	GetUserInProject(int) ([]UserResponse, error)
	CreateProject(NewProjectRequest) (*ProjectResponse, error)
	AddUserInProject(int, int) error
}
