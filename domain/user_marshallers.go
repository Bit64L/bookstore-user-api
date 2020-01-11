package domain

import (
	"bookstore-user-api/utils/errors"
	"encoding/json"
)

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

type Users []*User

func (users Users) Marshal(isPublic bool) ([]interface{}, *errors.RestErr) {
	results := make([]interface{}, len(users))
	for index, user := range users {
		marshaledUser, marshalErr := user.Marshal(isPublic)
		if marshalErr != nil {
			return nil, marshalErr
		}
		results[index] = marshaledUser
	}
	return results, nil
}

func (user *User) Marshal(isPublic bool) (interface{}, *errors.RestErr) {

	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	if isPublic {
		var publicUser PublicUser
		if err := json.Unmarshal(userJson, &publicUser); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		return publicUser, nil
	} else {
		var privateUser PrivateUser
		if err := json.Unmarshal(userJson, &privateUser); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		return privateUser, nil
	}
}
