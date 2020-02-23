package user

import "usermanagement/models"

type UserRepository interface {
	CheckUserEmail(email string) (bool, error)
	GetUserByEmail(email string) (*models.User, error)
	Register(user *models.User) (*models.User, error)
	GetAllUser() ([]*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(id uint) (bool, error)
}
