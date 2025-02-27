package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/your-username/RestApiGo/internal/database"
	"github.com/your-username/RestApiGo/internal/handlers"
	"github.com/your-username/RestApiGo/internal/taskService"
	"github.com/your-username/RestApiGo/internal/web/tasks"
	"log"
)

func main() {
	database.InitDB()
	//database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)

	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}
