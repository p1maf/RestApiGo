package userService

import "github.com/your-username/RestApiGo/internal/web/users"

type UserService struct {
	Repository UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{Repository: r}
}

func (service *UserService) CreateUser(user users.NewUser) (users.User, error) {
	return service.Repository.CreateUser(user)
}

func (service *UserService) GetAllUsers() ([]users.User, error) {
	return service.Repository.GetAllUsers()
}

func (service *UserService) DeleteUserById(id uint) error {
	return service.Repository.DeleteUserById(id)
}

func (service *UserService) UpdateUserById(id uint, task users.UpdateUser) (users.User, error) {
	return service.Repository.UpdateUserById(id, task)
}
