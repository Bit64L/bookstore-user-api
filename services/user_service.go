package services

import (
	"bookstore-user-api/domain"
	"bookstore-user-api/utils/errors"
)

type UserServiceInterface interface {
	CreateUser(domain.User) (*domain.User, *errors.RestErr)
	GetUser(int64) (*domain.User, *errors.RestErr)
	UpdateUser(domain.User) (*domain.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	FindUserByStatus(string) (domain.Users, *errors.RestErr)
}

var (
	UserService UserServiceInterface = &userService{}
)

type userService struct{}

func (userService *userService) CreateUser(user domain.User) (*domain.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (userService *userService) GetUser(userId int64) (*domain.User, *errors.RestErr) {
	result := &domain.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (userService *userService) UpdateUser(user domain.User) (*domain.User, *errors.RestErr) {
	if err := user.Update(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (userService *userService) DeleteUser(userId int64) *errors.RestErr {
	result := &domain.User{Id: userId}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil

}

func (userService *userService) FindUserByStatus(status string) (domain.Users, *errors.RestErr) {
	result := &domain.User{Status: status}
	users, err := result.FindByStatus()
	if err != nil {
		return nil, err
	}
	return users, nil
}
