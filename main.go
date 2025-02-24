package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var task []Task

	if err := DB.Find(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := DB.Create(&task).Error; err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	result := DB.Model(&Task{}).Where("id = ?", task.ID).Updates(task)

	if result.RowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	var updatedTask Task
	if err := DB.First(&updatedTask, task.ID).Error; err != nil {
		http.Error(w, "Failed to retrieve updated task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "No ID in request", http.StatusBadRequest)
	}

	result := DB.Unscoped().Delete(&Task{}, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", GetTasks).Methods(http.MethodGet)
	router.HandleFunc("/api/tasks", CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/api/tasks", UpdateTask).Methods(http.MethodPut)
	router.HandleFunc("/api/tasks/{id}", DeleteTask).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
