package repository

import (
	"a21hc3NpZ25tZW50/model" 

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var user []model.UserTaskCategory
	err := r.db.
		Table("users").
		Select("users.id, users.fullname, users.email, tasks.title AS task, tasks.deadline, tasks.priority, tasks.status, categories.name AS category").
		Joins("JOIN tasks ON users.id = tasks.user_id ").
		Joins("JOIN categories ON tasks.category_id = categories.id ").
		Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil // TODO: replace this
}
