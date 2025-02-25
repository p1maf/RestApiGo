package taskService

type TaskService struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repository.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repository.GetAllTasks()
}

func (s *TaskService) UpdateTaskById(id uint, task Task) (Task, error) {
	return s.repository.UpdateTaskById(id, task)
}

func (s *TaskService) DeleteTaskById(id string) error {
	return s.repository.DeleteTaskById(id)
}
