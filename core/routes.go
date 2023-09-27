package core

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (api *Api) Routes() {
	r := api.App.Group("/api/v1")

	{
		r.GET("/ping", ping)
	}
}
