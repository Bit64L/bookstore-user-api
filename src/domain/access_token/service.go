package access_token

import (
	"bookstore-oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpires(AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpires(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(accessToken AccessToken) *errors.RestErr {
	if err := s.repository.Create(accessToken); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateExpires(at AccessToken) *errors.RestErr {
	if err := s.repository.UpdateExpires(at); err != nil {
		return err
	}
	return nil
}
