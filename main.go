package main

import (
	"fmt"
	"todo/configs"
	"todo/controllers"
	"todo/repositories"
	"todo/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	configs.InitConfig()
	configs.InitTimeZone()
	db = configs.InitDatabase()
}

func main() {
	app := fiber.New()

	userRepository := repositories.NewUserRepositoryDB(db)
	userProjectRepository := repositories.NewUserProjectRepositoryDB(db)
	projectRepository := repositories.NewProjectRepositoryDB(db)
	todoRepository := repositories.NewTodoRepositoryDB(db)

	userService := services.NewUserService(userRepository)
	projectService := services.NewProjectService(userRepository, projectRepository, userProjectRepository)
	todoService := services.NewTodoService(todoRepository, userRepository, projectRepository)

	userController := controllers.NewUserController(userService)
	projectController := controllers.NewProjectController(projectService)
	todoController := controllers.NewTodoController(todoService)

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	api := app.Group("/api/v2")

	user := api.Group("/user")
	user.Get("/", userController.GetUsers)
	user.Get("/:id", userController.GetOneUser)
	user.Post("/", userController.CreateUser)

	project := api.Group("/project")
	project.Get("/", projectController.GetProjects)
	project.Post("/", projectController.CreateProject)
	project.Post("/adduser", projectController.AddUserInProject)

	todo := api.Group("/todo")
	todo.Post("/", todoController.CreateTodo)
	todo.Post("/", todoController.UpdateTodoStep)

	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
