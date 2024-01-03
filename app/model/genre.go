package model

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Genre struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index,priority:1; type:varchar(128)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (r *Genre) Create() (error) {
	err := utils.DB.Create(&r).Error

	return  err
}

func (r *Genre) GetAll() ([]*Genre, error) {
	genres := []*Genre{}
	err := utils.DB.Find(&genres).Error

	return genres, err
}

func (r *Genre) GetByID(id uint) (*Genre, error) {
	err := utils.DB.First(&r).Error

	return r, err
}