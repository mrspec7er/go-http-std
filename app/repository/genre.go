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

type Genres []*Genre

func (Genre) Create(req *Genre) (error) {
	err := utils.DB.Create(&req).Error

	return  err
}

func (Genres) GetAll() (*Genres, error) {
	m := &Genres{}
	err := utils.DB.Find(&m).Error

	return m, err
}

func (Genre) GetByID(id uint) (*Genre, error) {
	m := &Genre{ID: id}
	err := utils.DB.First(&m).Error

	return m, err
}