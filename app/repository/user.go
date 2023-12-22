package repository

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index,priority:1; type:varchar(128)"`
	Email string `json:"email" gorm:"unique;type:varchar(128)"`
	Password string `json:"password" gorm:"type:varchar(128)"`
	Status string `json:"status" gorm:"type:varchar(32)"`
	Role string `json:"role" gorm:"type:varchar(32)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}