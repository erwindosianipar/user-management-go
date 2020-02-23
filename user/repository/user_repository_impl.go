package repository

import (
	"errors"

	"usermanagement/models"
	"usermanagement/user"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var _ user.UserRepository = &UserRepositoryImpl{}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) CheckUserEmail(email string) (bool, error) {
	var total int

	if err := u.db.Table("users").Where("email = ?", email).Count(&total).Error; err != nil {
		logrus.Error("[UserRepositoryImpl.CheckUserEmail] ", err)
		return false, errors.New("ERROR when check user email")
	}

	if total > 0 {
		return true, nil
	}

	return false, nil
}

func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	dataUser := new(models.User)

	if err := u.db.Table("users").Where("email = ?", email).First(&dataUser).Error; err != nil {
		logrus.Error("[UserRepositoryImpl.GetUserByEmail] ", err)
		return nil, errors.New("ERROR when get user by email")
	}

	return dataUser, nil
}

func (u *UserRepositoryImpl) Register(user *models.User) (*models.User, error) {
	if err := u.db.Table("users").Save(&user).Error; err != nil {
		logrus.Error("[UserRepositoryImpl.Register] ", err)
		return nil, errors.New("ERROR when insert data user")
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetAllUser() ([]*models.User, error) {
	listDataUser := make([]*models.User, 0)

	if err := u.db.Table("users").Find(&listDataUser).Error; err != nil {
		logrus.Error("[UserRepositoryImpl.GetAllUser] ", err)
		return nil, errors.New("ERROR when get all data user")
	}

	return listDataUser, nil
}

func (u *UserRepositoryImpl) GetUserByID(id uint) (*models.User, error) {
	dataUser := new(models.User)

	if err := u.db.Table("users").Where("id = ?", id).Find(&dataUser).Error; err != nil {
		logrus.Error("[UserRepositoryImpl.GetUserByID] ", err)
		return nil, errors.New("ERROR when get user by id")
	}

	return dataUser, nil
}

func (u *UserRepositoryImpl) UpdateUser(id uint, user *models.User) (*models.User, error) {
	dataUser := new(models.User)

	if err := u.db.Table("users").Where("id = ?", id).First(&dataUser).Update(&user).Error; err != nil {
		logrus.Error("[UserRepositoryImpl.UpdateUser] ", err)
		return nil, errors.New("ERROR when update data user")
	}

	return dataUser, nil
}

func (u *UserRepositoryImpl) DeleteUser(id uint) (bool, error) {
	if err := u.db.Table("users").Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		logrus.Error("[UserRepositoryImpl.DeleteUser] ", err)
		return false, errors.New("ERROR when delete data user")
	}

	return true, nil
}

func CreateUserRepositoryImpl(db *gorm.DB) user.UserRepository {
	return &UserRepositoryImpl{db}
}
