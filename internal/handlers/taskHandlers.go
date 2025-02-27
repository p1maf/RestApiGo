package handlers

import (
	"context"
	"github.com/your-username/RestApiGo/internal/taskService"
	"github.com/your-username/RestApiGo/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) PatchTasks(_ context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {

	taskToUpdate := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskById(*request.Body.Id, taskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasks200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {

	err := h.Service.DeleteTaskById(request.Id)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, err
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

//func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
//	tasks, err := h.Service.GetAllTasks()
//	if err != nil {
//		http.Error(w, "Tasks not found", http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(tasks)
//}
//
//func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
//	var task taskService.Task
//
//	err := json.NewDecoder(r.Body).Decode(&task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	createdTask, err := h.Service.CreateTask(task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(createdTask)
//}
//
//func (h *Handler) PutTaskHandler(w http.ResponseWriter, r *http.Request) {
//	var task taskService.Task
//	err := json.NewDecoder(r.Body).Decode(&task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	updatedTask, err := h.Service.UpdateTaskById(task.ID, task)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(updatedTask)
//}
//
//func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id, ok := vars["id"]
//	if !ok {
//		http.Error(w, "No ID in request", http.StatusBadRequest)
//	}
//
//	err := h.Service.DeleteTaskById(id)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//}
