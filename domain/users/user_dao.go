package users

import (
	"github.com/amit-dhakad/bookstore_user-api/datasources/mysql/usersdb"
	"github.com/amit-dhakad/bookstore_user-api/utils/dateutils"
	"github.com/amit-dhakad/bookstore_user-api/utils/errors"
	"github.com/amit-dhakad/bookstore_user-api/utils/mysqlutils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id, first_name,last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?,last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

var (
	usersDB = make(map[int64]*User)
)

// Get get userby id
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())

	}

	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}
	return nil
}

// Save save user
func (user *User) Save() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {

		panic(err)
	}
	stmt, err := usersdb.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)

	}

	userID, err := insertResult.LastInsertId()

	if err != nil {
		return mysqlutils.ParseError(saveErr)
	}
	user.ID = userID
	return nil
}

// Update user in to database
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

// Delete error
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}
