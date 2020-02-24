package user

import "usermanagement/models"

type UserService interface {
	CheckUserEmail(email string) (bool, error)
	GetUserByEmail(email string) (*models.User, error)
	Register(user *models.User) (map[string]interface{}, error)
	GetAllUser() (map[string]interface{}, error)
	GetUserByID(id uint) (map[string]interface{}, error)
	UpdateUser(id uint, user *models.User) (map[string]interface{}, error)
	DeleteUser(id uint) (map[string]interface{}, error)
}
