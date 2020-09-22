package services

import (
	"github.com/amit-dhakad/bookstore_user-api/domain/users"
	"github.com/amit-dhakad/bookstore_user-api/utils/errors"
)

// CreateUser create user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil
}
