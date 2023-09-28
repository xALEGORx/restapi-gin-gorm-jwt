package middlewares

import (
	"fmt"
	"os"
	"restapi/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Handler(db *gorm.DB, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Set("log", log)

		// Parse jwt token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Next()
			return
		}

		bearer_token := strings.Split(token, "Bearer ")
		if len(bearer_token) < 1 {
			c.Next()
			return
		}

		token = bearer_token[1]
		tokenData, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT")), nil
		})
		if err != nil || !tokenData.Valid {
			c.Next()
			return
		}

		userId := tokenData.Claims.(jwt.MapClaims)["user"]
		user := models.User{}
		if err := db.First(&user, userId).Error; err != nil {
			c.Next()
			return
		}
		c.Set("user", user)

		// Finish handler
		c.Next()
	}
}
