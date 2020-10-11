package users

import (
	"fmt"
	"strings"

	"github.com/amit-dhakad/bookstore_user-api/datasources/mysql/usersdb"
	"github.com/amit-dhakad/bookstore_user-api/utils/dateutils"
	"github.com/amit-dhakad/bookstore_user-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name,last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

// Get get userby id
func (user *User) Get() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
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
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exits", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save users: %s", err.Error()))
	}

	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))

	}
	user.ID = userID
	return nil
}
