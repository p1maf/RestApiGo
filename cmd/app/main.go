package main

import (
	"github.com/gorilla/mux"
	"github.com/your-username/RestApiGo/internal/database"
	"github.com/your-username/RestApiGo/internal/handlers"
	"github.com/your-username/RestApiGo/internal/taskService"
	"log"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetTasksHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/tasks", handler.PutTaskHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
