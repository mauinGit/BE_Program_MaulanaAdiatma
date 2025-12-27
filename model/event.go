package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Judul	        string		`json:"judul" gorm:"not null"`
	Cover           string		`json:"cover" gorm:"not null"`
	Deskripsi       string		`json:"deskripsi" gorm:"not null"`
	Tanggal	        time.Time	`json:"tanggal" gorm:"not null"`
	Capacity        int			`json:"capacity" gorm:"not null"`
	RemainingTicket int 		`json:"remaining_ticket" gorm:"default:0"`
	Price           int			`json:"price" gorm:"not null"`
}
