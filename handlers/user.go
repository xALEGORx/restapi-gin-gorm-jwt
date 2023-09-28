package handlers

import (
	"restapi/models"
	"restapi/types"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	c.JSON(200, types.RESPONSE{
		Success: true,
		Data:    user.PrepareToView(),
	})
}
