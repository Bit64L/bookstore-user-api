package domain

import (
	"bookstore-user-api/util/errors"
	"fmt"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s has been registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d exists", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
