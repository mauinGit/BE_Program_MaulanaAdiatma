package model

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID  uint `json:"user_id" gorm:"not null"`
	EventID uint `json:"event_id" gorm:"not null"`

	Quantity   int `json:"quantity" gorm:"not null"`
	TotalPrice int `json:"total_price" gorm:"not null"`
}

