package users

import (
	"net/http"
	"strconv"

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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number ")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number ")
		c.JSON(err.Status, err)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(
			"invalid json body",
		)
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)

}
