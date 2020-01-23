package domain

import (
	"bookstore-user-api/datasource/mysql/user_db"
	"bookstore-user-api/logger"
	"bookstore-user-api/utils/constant"
	"bookstore-user-api/utils/crypto_utils"
	"bookstore-user-api/utils/date_utils"
	"bookstore-user-api/utils/errors"
)

const (
	queryInsertUser                 = "INSERT INTO user(first_name, last_name, email, date_created, password, status) values(?,?,?,?,?,?);"
	queryGetUser                    = "select id, first_name, last_name, date_created, email, password, status from user where id = ?;"
	queryUpdateUser                 = "update user set first_name = ?, last_name = ?, email = ?, password = ?, status = ? where id = ?;"
	queryDeleteUser                 = "delete from user where id = ?"
	queryFindUserByStatus           = "select id, first_name, last_name, date_created, email, password, status from user where status = ?"
	queryFindUserByEmailAndPassword = "select id, first_name, last_name, date_created, email, password, status from user where email = ? and password = ?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Email, &user.Password, &user.Status); err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	user.Status = constant.USER_ACTIVE
	user.Password = crypto_utils.GetSha256(user.Password)

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	user.Id, err = result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get user's id", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	return nil
}

func (user *User) FindByStatus() ([]*User, *errors.RestErr) {
	stmt, err := user_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user statement", err)
		return nil, errors.NewInternalServerError(errors.DatabaseError)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Status)
	if err != nil {
		logger.Error("error when trying to query user", err)
		return nil, errors.NewInternalServerError(errors.DatabaseError)
	}

	users := make([]*User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Email, &user.Password, &user.Status); err != nil {
			logger.Error("error when trying to get user", err)
			return nil, errors.NewInternalServerError(errors.DatabaseError)
		}
		users = append(users, &user)
	}

	return users, nil

}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := user_db.Client.Prepare(queryFindUserByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare find user by email and password statement", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}
	defer stmt.Close()

	rows := stmt.QueryRow(user.Email, crypto_utils.GetSha256(user.Password))
	if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.Email, &user.Password, &user.Status); err != nil {
		logger.Error("error when trying to scan user", err)
		return errors.NewInternalServerError(errors.DatabaseError)
	}

	return nil

}
