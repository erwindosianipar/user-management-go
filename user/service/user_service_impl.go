package service

import (
	"usermanagement/commons"
	"usermanagement/models"
	"usermanagement/user"
)

type UserService struct {
	userRepository user.UserRepository
}

func (u *UserService) CheckUserEmail(email string) (bool, error) {
	status, err := u.userRepository.CheckUserEmail(email)
	if err != nil {
		return status, err
	}

	return status, nil
}

func (u *UserService) GetUserByEmail(email string) (*models.User, error) {
	response, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (u *UserService) Register(user *models.User) (map[string]interface{}, error) {
	isEmailExist, err := u.userRepository.CheckUserEmail(user.Email)
	if err != nil {
		return commons.Message(false, err.Error()), err
	}

	if isEmailExist {
		return commons.Message(false, "error: email already exist"), err
	}

	response, err := u.userRepository.Register(user)
	if err != nil {
		return commons.Message(false, err.Error()), err
	}

	mapResponse := commons.Message(true, "success: register user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) GetAllUser() (map[string]interface{}, error) {
	response, err := u.userRepository.GetAllUser()
	if err != nil {
		return commons.Message(false, err.Error()), err
	}

	mapResponse := commons.Message(true, "success: get all data user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) GetUserByID(id uint) (map[string]interface{}, error) {
	response, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return commons.Message(false, err.Error()), err
	}

	mapResponse := commons.Message(true, "success: get data user by id")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) UpdateUser(id uint, user *models.User) (map[string]interface{}, error) {
	userData, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return commons.Message(false, err.Error()), nil
	}

	isEmailExist, err := u.userRepository.CheckUserEmail(user.Email)
	if err != nil {
		return commons.Message(false, err.Error()), nil
	}

	if isEmailExist {
		if user.Email != userData.Email {
			return commons.Message(false, "error: email already used"), nil
		}
	}

	response, err := u.userRepository.UpdateUser(id, user)
	if err != nil {
		return commons.Message(false, err.Error()), nil
	}

	mapResponse := commons.Message(true, "success: update data user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) DeleteUser(id uint) (map[string]interface{}, error) {
	_, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return commons.Message(false, err.Error()), nil
	}

	response, err := u.userRepository.DeleteUser(id)
	if err != nil {
		return commons.Message(false, err.Error()), err
	}

	mapResponse := commons.Message(true, "success: delete data user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func CreateUserService(userRepository user.UserRepository) user.UserService {
	return &UserService{userRepository}
}
