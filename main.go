package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var task string

type TaskResponse struct {
	Task string `json:"task"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello", task)
}

func SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taskResponse TaskResponse

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&taskResponse); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	task = taskResponse.Task
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/setTask", SetTaskHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
