package model

import (
	"strings"
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Movie struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"index,priority:1; type:varchar(128)"`
	Thumbnail string `json:"thumbnail" gorm:"type:text"`
	Description string `json:"description" gorm:"type:text"`
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
	// Belongs to relation
	GenreID uint `json:"genreId"`
	Genre *Genre `json:"genre" gorm:"constraint:OnDelete:SET NULL"`
	// Has many relation
	DirectorID uint `json:"directorId"`
	Director *Director `json:"director" gorm:"constraint:OnDelete:SET NULL"`
	// Many 2 many relation
	Casts []*Cast `json:"casts" gorm:"many2many:cast_movies;"`
	// One 2 many relation
	Photos []*Photo `json:"photos" gorm:"constraint:OnDelete:SET NULL"`
}

func (r *Movie) Create() (error) {
	err := utils.DB.Create(&r).Error

	return  err
}

func (r *Movie) GetAll(offset int, limit int, keyword string) ([]*Movie, *int64, error) {
	movies := []*Movie{}
	var count int64

	query := utils.DB.Offset(offset * limit).Limit(limit)
	if keyword != "" {
		query = query.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ?", "%"+strings.ToLower(keyword)+"%", "%"+strings.ToLower(keyword)+"%")
	}
	err := query.Preload("Genre").Preload("Director").Preload("Casts").Find(&movies).Offset(-1).Count(&count).Error

	return movies, &count, err
}

func (r *Movie) GetByID(id uint) (*Movie, error) {
	r.ID = id
	err := utils.DB.Preload("Genre").Preload("Director").Preload("Casts").First(&r).Error

	return r, err
}

func (r *Movie) Update() (error) {
	err := utils.DB.Updates(&r).Error
	return  err
}

func (r *Movie) UpdateModelAndAssociation() (error) {
	err := utils.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&r).Error
	return  err
}

func (r *Movie) UpdateAssociation() (error) {
	res := utils.DB.Model(&r).Association("Casts").Replace(r.Casts)
	return  res
}