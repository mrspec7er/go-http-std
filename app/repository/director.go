package repository

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Director struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index,priority:1; type:varchar(128)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type Directors []*Director

func (Director) Create(req *Director) (error) {
	err := utils.DB.Create(&req).Error

	return  err
}

func (Directors) GetAll() (*Directors, error) {
	m := &Directors{}
	err := utils.DB.Find(&m).Error

	return m, err
}

func (Director) GetByID(id uint) (*Director, error) {
	m := &Director{ID: id}
	err := utils.DB.First(&m).Error

	return m, err
}