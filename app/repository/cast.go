package repository

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Cast struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index,priority:1; type:varchar(128)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	// Many 2 many relation
	Movies *[]Movie `json:"movies" gorm:"many2many:cast_movies;"`
}

type Casts []*Cast

func (Cast) Create(req *Cast) (error) {
	err := utils.DB.Create(&req).Error

	return  err
}

func (Casts) GetAll() (*Casts, error) {
	m := &Casts{}
	err := utils.DB.Preload("Movies").Find(&m).Error

	return m, err
}

func (Cast) GetByID(id uint) (*Cast, error) {
	m := &Cast{ID: id}
	err := utils.DB.Preload("Movies").First(&m).Error

	return m, err
}