package repository

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

func (g *Genre) Create() (error) {
	err := utils.DB.Create(&g).Error

	return  err
}

func (g *Genre) GetAll() ([]*Genre, error) {
	genres := []*Genre{}
	err := utils.DB.Find(&genres).Error

	return genres, err
}

func (g *Genre) GetByID(id uint) (*Genre, error) {
	err := utils.DB.First(&g).Error

	return g, err
}