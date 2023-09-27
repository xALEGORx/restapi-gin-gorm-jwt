package handlers

import (
	"restapi/types"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, types.RESPONSE{
		Success: true,
		Data: types.JSON{
			"message": "pong",
		},
	})
}
