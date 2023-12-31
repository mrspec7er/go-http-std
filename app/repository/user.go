package repository

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (u *User) Create() (error) {
	err := utils.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&u).Error

	return  err
}

func (u *User) GetAll() ([]*User, error) {
	users := []*User{}
	err := utils.DB.Find(&users).Error

	return users, err
}

func (u *User) GetByID(id uint) (*User, error) {
	u.ID = id
	err := utils.DB.First(&u).Error

	return u, err
}