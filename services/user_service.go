package services

import (
	"net/mail"
	"todo/repositories"
	"todo/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusInternalServerError,
			Message: "internal error",
		}
	}

	usersData := []UserResponse{}
	for _, value := range users {
		usersData = append(usersData, UserResponse{
			ID:        value.ID,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Email:     value.Email,
			Active:    value.Active,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return usersData, nil
}

func (s userService) GetOneUser(id uint) (*UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusNotFound,
			Message: "user id not found",
		}
	}

	userData := UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &userData, nil
}

func (s userService) Login(email, password string) (*TokenResponse, error) {
	return &TokenResponse{}, nil
}

func (s userService) CreateUser(newUser NewUserRequest) (*UserResponse, error) {

	//validate email
	_, err := mail.ParseAddress(newUser.Email)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusBadRequest,
			Message: "email is invalid",
		}
	}

	//check have this email yet?
	userEmail, _ := s.userRepo.GetByEmail(newUser.Email)
	if userEmail != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusConflict,
			Message: "email already exists",
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "encryption password is failed",
		}
	}

	//create user
	addUser := repositories.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  string(password),
	}

	user, err := s.userRepo.Create(addUser)
	if err != nil {
		return nil, utils.HandlerErr{
			Code:    fiber.StatusInternalServerError,
			Message: "failed creation",
		}
	}

	resUser := UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &resUser, nil
}
