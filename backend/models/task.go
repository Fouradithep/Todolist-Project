package models

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Status    string         `gorm:"default:Pending" json:"status"` // Pending, Doing , Done
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID 	  uint   // FK
  	User      User   // สร้างความสัมพันธ์ของUserID
}
