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
	// Has many relation
	Movies *[]Movie `json:"movies"`
}

func (d *Director) Create() (error) {
	err := utils.DB.Create(&d).Error

	return  err
}

func (d *Director) GetAll() ([]*Director, error) {
	directors := []*Director{}
	err := utils.DB.Preload("Movies").Find(&directors).Error

	return directors, err
}

func (d *Director) GetByID(id uint) (*Director, error) {
	d.ID = id
	err := utils.DB.Preload("Movies").First(&d).Error

	return d, err
}