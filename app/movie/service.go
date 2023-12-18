package movie

import (
	"fmt"

	"gorm.io/gorm"
)

type Movie struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Director string `json:"director"`
	Genres string `json:"genres"`
	Cast string `json:"cast"`
	DateAdded string `json:"dateAdded"`
}

type MovieService struct {
	DB *gorm.DB
}

func (m MovieService) Create(movie Movie) (int, error) {
	err := m.DB.Create(&movie).Error

	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (m MovieService) GetAll() {
	fmt.Println("Get all Movie")
}

func (m MovieService) GetOne(id uint) {
	fmt.Println("Get one Movie")
}

func (m MovieService) Update(movie Movie) {
	fmt.Println("Update a Movie")
}

func (m MovieService) Delete(id uint) {
	fmt.Println("Delete a Movie")
}