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
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name,last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id, first_name,last_name, email, date_created FROM users WHERE id=?;"
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
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)

		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d does not exits", user.ID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.ID, err.Error()))
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
