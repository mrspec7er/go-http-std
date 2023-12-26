package repository

import (
	"strings"
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Movie struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"index,priority:1; type:varchar(128)"`
	Description string `json:"description" gorm:"type:text"`
	Director string `json:"director" gorm:"type:varchar(128)"`
	Cast string `json:"cast" gorm:"index,priority:2; type:varchar(256)"`
	ProductionCountry string `json:"productionCountry" gorm:"type:varchar(128)"`
	ReleaseDate string `json:"releaseDate" gorm:"type:varchar(64)"`
	Rating string `json:"rating" gorm:"type:varchar(64)"`
	Duration string `json:"duration" gorm:"type:varchar(64)"`
	IMDBScore string `json:"imdbScore" gorm:"type:varchar(64)"`
	ContentType string `json:"contentType" gorm:"type:varchar(64)"`
	DateAdded string `json:"dateAdded" gorm:"type:varchar(64)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	// Has one relation
	GenreID uint `json:"genreId"`
	Genre *Genre `json:"genre" gorm:"constraint:OnDelete:SET NULL"`
}

type Movies []*Movie

func (Movie) Create(req *Movie) (error) {
	err := utils.DB.Create(&req).Error

	return  err
}

func (Movies) GetAll(offset int, limit int, keyword string) (*Movies, *int64, error) {
	m := &Movies{}
	var count int64

	query := utils.DB.Offset(offset * limit).Limit(limit)
	if keyword != "" {
		query = query.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", "%"+strings.ToLower(keyword)+"%", "%"+strings.ToLower(keyword)+"%")
	}
	err := query.Preload("Genre").Find(&m).Offset(-1).Count(&count).Error

	return m, &count, err
}

func (Movie) GetByID(id uint) (*Movie, error) {
	m := &Movie{ID: id}
	err := utils.DB.Preload("Genre").First(&m).Error

	return m, err
}