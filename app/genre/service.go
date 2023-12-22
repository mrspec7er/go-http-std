package genre

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type GenreService struct {}

func (GenreService) Create(req repository.Genre) (int, error) {
	movie := &repository.Genre{}

	err := movie.Create(&req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (GenreService) GetAll() ([]*repository.Genre, int, error) {
	movies := &repository.Genres{}

	result, err := movies.GetAll()

	if err != nil {
		return *result, 500, err
	}

	return *result, 200, nil
}

func (GenreService) GetOne(id uint) (*repository.Genre, int, error) {
	movie := &repository.Genre{}

	result, err := movie.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (GenreService) Update(movie repository.Genre) {
	fmt.Println("Update a Movie")
}

func (GenreService) Delete(id uint) {
	fmt.Println("Delete a Movie")
}