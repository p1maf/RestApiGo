package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/your-username/RestApiGo/internal/database"
	"github.com/your-username/RestApiGo/internal/handlers"
	"github.com/your-username/RestApiGo/internal/taskService"
	"github.com/your-username/RestApiGo/internal/userService"
	"github.com/your-username/RestApiGo/internal/web/tasks"
	"github.com/your-username/RestApiGo/internal/web/users"
	"log"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewUserService(usersRepo)
	usersHandler := handlers.UserNewHandler(usersService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(tasksHandler, nil)
	strictUserHandler := users.NewStrictHandler(usersHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}
