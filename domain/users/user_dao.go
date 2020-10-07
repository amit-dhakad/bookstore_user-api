package users

import (
	"fmt"

	"github.com/amit-dhakad/bookstore_user-api/utils/dateutils"
	"github.com/amit-dhakad/bookstore_user-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get get userby id
func (user *User) Get() *errors.RestErr {
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
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exits", user.ID))
	}

	user.DateCreated = dateutils.GetNowString()
	usersDB[user.ID] = user
	return nil
}
