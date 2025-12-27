package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Role     string `json:"role" gorm:"type:enum('admin','user');default:'user'"`
}
