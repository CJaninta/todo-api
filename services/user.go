package services

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"email,required"`
	Password  string `json:"password" validate:"required"`
}

type TokenResponse struct {
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetOneUser(uint) (*UserResponse, error)
	Login(string, string) (*TokenResponse, error)
	CreateUser(NewUserRequest) (*UserResponse, error)
}
