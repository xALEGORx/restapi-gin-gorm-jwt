package middlewares

import (
	"restapi/types"

	"github.com/gin-gonic/gin"
)

func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.Error(types.UNAUTHORIZED)
		c.Abort()
		return
	}
	c.Next()
}
