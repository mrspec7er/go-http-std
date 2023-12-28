package repository

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
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

type Users []*User

func (User) Create(req *User) (error) {
	err := utils.DB.Create(&User{Name: req.Name, Email: req.Email, Password: "Encrypted", Status: "INACTIVE", Role: "USER"}).Error

	return  err
}

func (Users) GetAll() (*Users, error) {
	m := &Users{}
	err := utils.DB.Find(&m).Error

	return m, err
}

func (User) GetByID(id uint) (*User, error) {
	m := &User{ID: id}
	err := utils.DB.First(&m).Error

	return m, err
}