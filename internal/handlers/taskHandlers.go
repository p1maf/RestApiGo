package handlers

import (
	"context"
	"github.com/your-username/RestApiGo/internal/taskService"
	"github.com/your-username/RestApiGo/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskToUpdate := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: *request.Body.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskById(request.Id, taskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h *Handler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	tasksUser, err := h.Service.GetTaskByUserId(request.Id)
	if err != nil {
		return nil, err
	}
	response := tasks.GetUsersIdTasks202JSONResponse{}

	for _, tsk := range tasksUser {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
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
		UserId: *taskRequest.UserId,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserId,
	}
	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
