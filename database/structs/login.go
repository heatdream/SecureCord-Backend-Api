package structs

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique; not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
	UUID     string `gorm:"unique"`
}
	