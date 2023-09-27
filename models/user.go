package models

import (
	"restapi/types"
	"time"
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
