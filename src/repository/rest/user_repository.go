package rest

import (
	"bookstore-oauth-api/src/domain/users"
	"bookstore-oauth-api/src/utils/errors"
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com:5001",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct{}

func NewRepository() RestUserRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {

	loginRequest := users.LoginRequest{
		Email:    email,
		Password: password,
	}

	response := userRestClient.Post("/internal/users/login", loginRequest)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("error when trying get response from login")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			return nil, errors.NewInternalServerError("error when trying to unmarshal rest error")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal user")
	}
	return &user, nil

}
