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

func (c *Cast) Create() (error) {
	err := utils.DB.Create(&c).Error

	return  err
}

func (c *Cast) GetAll() ([]*Cast, error) {
	casts := []*Cast{}
	err := utils.DB.Preload("Movies").Find(&casts).Error

	return casts, err
}

func (c *Cast) GetByID(id uint) (*Cast, error) {
	c.ID = id
	err := utils.DB.Preload("Movies").First(&c).Error

	return c, err
}