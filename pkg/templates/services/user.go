package services

import (
	"{{.Appname}}/models"
	"{{.Appname}}/repositories"
)

// UserService ...
type UserService struct {
	User *models.User
}

// NewUserService ...
func NewUserService() *UserService {
	return &UserService{
		User: models.NewUser(),
	}
}

// GetUserInfo ...
func (service *UserService) GetUserInfo(id uint) (resp models.User, err error) {
	resp, err = repositories.GetUserInfo(id)
	return
}

// GetUserInfo ...
func (service *UserService) GetUserList() (resp []models.User, err error) {
	resp, err = repositories.GetUserList()
	return
}

// GetUserInfo ...
func (service *UserService) CreateUser(username, password string) (err error) {

	user := models.NewUser()
	user.Username = username
	user.Password = password

	err = repositories.CreateUser(user)

	return
}

// GetUserInfo ...
func (service *UserService) UpdateUserInfo(id uint, password string) (err error) {
	item := models.User{}
	item.ID = id
	item.Password = password
	err = repositories.UpdateUserInfo(&item)
	return
}

// GetUserInfo ...
func (service *UserService) DeleteUser(id uint) (err error) {

	err = repositories.DeleteUser(id)
	return
}
