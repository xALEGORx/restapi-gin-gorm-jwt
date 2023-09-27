package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Set("log", log)

		c.Next()
	}
}
