package userService

import (
	"github.com/your-username/RestApiGo/internal/web/users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user users.NewUser) (users.User, error)
	GetAllUsers() ([]users.User, error)
	UpdateUserById(id uint, task users.UpdateUser) (users.User, error)
	DeleteUserById(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(newUser users.NewUser) (users.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return users.User{}, err
	}
	hashedPasswordStr := string(hashedPassword)
	user := users.User{
		Email:    &newUser.Email,
		Password: &hashedPasswordStr,
	}

	result := repo.db.Create(&user)
	if result.Error != nil {
		return users.User{}, result.Error
	}

	return user, nil
}

func (repo *userRepository) GetAllUsers() ([]users.User, error) {
	var users []users.User
	result := repo.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *userRepository) DeleteUserById(id uint) error {
	result := repo.db.Unscoped().Delete(&users.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *userRepository) UpdateUserById(id uint, user users.UpdateUser) (users.User, error) {

	result := repo.db.Model(&users.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return users.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return users.User{}, result.Error
	}

	if result.Error != nil {
		return users.User{}, result.Error
	}

	var updatedUser users.User
	if err := repo.db.First(&updatedUser, id).Error; err != nil {
		return users.User{}, err
	}

	return updatedUser, nil
}
