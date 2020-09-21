package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping it will return pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
