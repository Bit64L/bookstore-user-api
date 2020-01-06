package domain

import (
	"bookstore-user-api/datasource/mysql/user_db"
	"bookstore-user-api/util/date_utils"
	"bookstore-user-api/util/errors"
	"fmt"
)

const (
	queryInsertUser = "INSERT INTO user(first_name, last_name, email, date_created) values(?,?,?,?);"
	queryGetUser    = "select id, first_name, last_name, date_created, email from user where id = ?;"
	queryUpdateUser = "update user set first_name = ?, last_name = ?, email = ? where id = ?;"
)

var (
	userDB map[int64]*User
)

func (user *User) Get() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Email); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user: %d, %s", user.Id,
			err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id, err = result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to update user: %s", err.Error()))
	}
	return nil
}
