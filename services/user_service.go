package services

import (
	"bookstore-user-api/domain"
	"bookstore-user-api/utils/errors"
)

func CreateUser(user domain.User) (*domain.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*domain.User, *errors.RestErr) {
	result := &domain.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user domain.User) (*domain.User, *errors.RestErr) {
	if err := user.Update(); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	result := &domain.User{Id: userId}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil

}

func FindUserByStatus(status string) (domain.Users, *errors.RestErr) {
	result := &domain.User{Status: status}
	users, err := result.FindByStatus()
	if err != nil {
		return nil, err
	}
	return users, nil
}
