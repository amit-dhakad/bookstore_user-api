package users

import (
	"net/http"
	"strconv"

	"github.com/amit-dhakad/bookstore_user-api/domain/users"
	"github.com/amit-dhakad/bookstore_user-api/services"
	"github.com/amit-dhakad/bookstore_user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number ")

	}
	return userID, nil
}

// Create will create user
func Create(c *gin.Context) {
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

// Get will return user
func Get(c *gin.Context) {
	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update update user to database
func Update(c *gin.Context) {

	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
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

	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)

}

// Delete user from database
func Delete(c *gin.Context) {

	userID, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search search user by status
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.Search(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
