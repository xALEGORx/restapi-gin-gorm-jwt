package handlers

import (
	"restapi/models"
	"restapi/types"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	body := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.BindJSON(&body); err != nil {
		c.Error(types.WRONG_PARAMETERS)
		return
	}

	user := models.User{}
	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.Error(types.USER_NOT_FOUND)
		return
	}

	if !user.CheckPassword(body.Password) {
		c.Error(types.WRONG_PASSWORD)
		return
	}

	token, err := user.GenerateToken()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, types.RESPONSE{
		Success: true,
		Data: types.JSON{
			"user":  user.PrepareToView(),
			"token": token,
		},
	})
}

func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	body := struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.BindJSON(&body); err != nil {
		c.Error(types.WRONG_PARAMETERS)
		return
	}

	exists := models.User{}
	if err := db.Where("email = ?", body.Email).First(&exists).Error; err == nil {
		c.Error(types.USER_EXIST)
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		c.Error(err)
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(bytes),
	}

	db.Create(&user)

	c.JSON(200, types.RESPONSE{
		Success: true,
		Data:    user.PrepareToView(),
	})
}
