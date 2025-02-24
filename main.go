package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
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
}

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", GetTask).Methods("GET")
	router.HandleFunc("/api/setTask", CreateTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
