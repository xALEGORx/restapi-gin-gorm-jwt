package models

import (
	"os"
	"restapi/types"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"user_id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) PrepareToView() types.JSON {
	return types.JSON{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
	}
}

func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) GenerateToken() (string, error) {
	expire, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	date := time.Now().Add(time.Second * time.Duration(expire))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": u.PrepareToView(),
		"exp":  date.Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	return tokenString, err
}
