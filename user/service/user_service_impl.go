package service

import (
	"usermanagement/common"
	"usermanagement/models"
	"usermanagement/user"
)

type UserService struct {
	userRepository user.UserRepository
}

func (u *UserService) CheckUserEmail(email string) (map[string]interface{}, error) {
	response, err := u.userRepository.CheckUserEmail(email)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "SUCCESS check user email")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) GetUserByEmail(email string) (map[string]interface{}, error) {
	response, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "SUCCESS get data user by email")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) Register(user *models.User) (map[string]interface{}, error) {
	response, err := u.userRepository.Register(user)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "SUCCESS register user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) GetAllUser() (map[string]interface{}, error) {
	response, err := u.userRepository.GetAllUser()
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "SUCCESS get all data user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) GetUserByID(id uint) (map[string]interface{}, error) {
	response, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "SUCCESS get data user by id")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) UpdateUser(id uint, user *models.User) (map[string]interface{}, error) {
	response, err := u.userRepository.UpdateUser(id, user)
	if err != nil {
		return common.Message(false, err.Error()), nil
	}

	mapResponse := common.Message(true, "SUCCESS update data user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (u *UserService) DeleteUser(id uint) (map[string]interface{}, error) {
	response, err := u.userRepository.DeleteUser(id)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "SUCCESS delete data user")
	mapResponse["response"] = response
	return mapResponse, nil
}

func CreateUserService(userRepository user.UserRepository) user.UserService {
	return &UserService{userRepository}
}
