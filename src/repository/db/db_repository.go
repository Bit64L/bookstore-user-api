package db

import (
	"bookstore-oauth-api/src/clients/cassandra"
	"bookstore-oauth-api/src/domain/access_token"
	"bookstore-oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "select access_token, user_id, client_id, expires from access_tokens where access_token = ?"
	queryCreateAccessToken = "insert into access_tokens(access_token, user_id, client_id, expires) values (?,?,?,?) "
	queryUpdateExpires     = "update access_tokens set expires = ? where access_token = ?"
)

func New() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpires(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	var accessToken access_token.AccessToken
	if scanErr := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&accessToken.AccessToken,
		&accessToken.UserId,
		&accessToken.ClientId,
		&accessToken.Expires); scanErr != nil {
		return nil, errors.NewInternalServerError(scanErr.Error())
	}

	return &accessToken, nil

}

func (r *dbRepository) Create(accessToken access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		accessToken.AccessToken,
		accessToken.UserId,
		accessToken.ClientId,
		accessToken.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpires(accessToken access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		accessToken.Expires,
		accessToken.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}
