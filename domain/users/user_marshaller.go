package users

import "encoding/json"

// PublicUser return public data
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser return private data except password
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall return user by given condition
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	// return PrivateUser{
	// 	ID:          user.ID,
	// 	FirstName:   user.FirstName,
	// 	LastName:    user.LastName,
	// 	Email:       user.Email,
	// 	DateCreated: user.DateCreated,
	// 	Status:      user.Status,
	// }

	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}
