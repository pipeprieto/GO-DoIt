package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	ID          int 		`gorm:"primaryKey"`
	Title       string 		`gorm:"type:varchar(100);not null; unique_index"`
	Done        bool  		`gorm:"default:false"`
	Description string		`gorm:"type:varchar(400)"`
	CreatedAt   time.Time 	 
}