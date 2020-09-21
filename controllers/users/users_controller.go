package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser will create user
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

// GetUser will return user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")

}
