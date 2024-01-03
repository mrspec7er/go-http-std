package repository

import (
	"time"

	"github.com/mrspec7er/go-http-std/app/utils"
	"gorm.io/gorm"
)

type Photo struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index,priority:1; type:varchar(128)"`
	URL string `json:"url" gorm:"type:text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	// One 2 many relation
	MovieID uint `json:"movieId"`
	Movie *Movie `json:"movie"`
}

func (r *Photo) BulkCreate(photos []*Photo) (error) {
	err := utils.DB.Create(&photos).Error

	return  err
}

func (r *Photo) GetByMovieId(movieId uint) ([]*Photo, error) {
	photos := []*Photo{}
	err := utils.DB.Where("movie_id = ?", movieId).Find(&photos).Error

	return photos, err
}

func (r *Photo) GetByID(id uint) (*Photo, error) {
	r.ID = id
	err := utils.DB.Preload("Movies").First(&r).Error

	return r, err
}

func (r *Photo) DeleteByID(id uint) (*Photo, error) {
	r.ID = id
	err := utils.DB.Delete(&r).Error

	return r, err
}

func (r *Photo) DeleteByMovie(movieId uint) (error) {
	err := utils.DB.Where("movie_id = ?", movieId).Delete(&r).Error

	return err
}