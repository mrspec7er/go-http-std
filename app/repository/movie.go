package repository

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Movie struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"index,priority:1; type:varchar(128)"`
	Description string `json:"description" gorm:"type:text"`
	Director string `json:"director" gorm:"type:varchar(128)"`
	Genres string `json:"genres" gorm:"index,priority:3; type:varchar(128)"`
	Cast string `json:"cast" gorm:"index,priority:2; type:varchar(256)"`
	ProductionCountry string `json:"productionCountry" gorm:"type:varchar(128)"`
	ReleaseDate string `json:"releaseDate" gorm:"type:varchar(64)"`
	Rating string `json:"rating" gorm:"type:varchar(64)"`
	Duration string `json:"duration" gorm:"type:varchar(64)"`
	ImdbScore string `json:"imdbScore" gorm:"type:varchar(64)"`
	ContentType string `json:"contentType" gorm:"type:varchar(64)"`
	DateAdded string `json:"dateAdded" gorm:"type:varchar(64)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type Movies []*Movie

func (Movie) Create(req *Movie) (error) {
	err := utils.DB.Create(&req).Error

	return  err
}

func (Movies) GetAll() (*Movies, error) {
	m := &Movies{}
	err := utils.DB.Find(&m).Error

	return m, err
}

func (Movie) GetByID(id uint) (*Movie, error) {
	m := &Movie{ID: id}
	err := utils.DB.First(&m).Error

	return m, err
}