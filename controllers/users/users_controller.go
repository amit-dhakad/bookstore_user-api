package users

import (
	"net/http"

	"github.com/amit-dhakad/bookstore_user-api/domain/users"
	"github.com/amit-dhakad/bookstore_user-api/services"
	"github.com/amit-dhakad/bookstore_user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser will create user
func CreateUser(c *gin.Context) {
	var user users.User
	/* bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		//TODO: Handle error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		// TODO: Handle json error
		return
	} */

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(
			"invalid json body",
		)
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)

	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser will return user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
