package model

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Cast struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"index,priority:1; type:varchar(128)"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	// Many 2 many relation
	Movies []*Movie `json:"movies" gorm:"many2many:cast_movies;"`
}

func (r *Cast) Create() error {
	err := utils.DB.Create(&r).Error
	return err
}

func (r *Cast) GetAll() ([]*Cast, error) {
	casts := []*Cast{}
	err := utils.DB.Preload("Movies").Find(&casts).Error
	return casts, err
}

func (r *Cast) GetByID(id uint) (*Cast, error) {
	r.ID = id
	err := utils.DB.Preload("Movies").First(&r).Error
	return r, err
}
