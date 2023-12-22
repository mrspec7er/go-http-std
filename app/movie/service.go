package movie

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type MovieService struct {}

func (MovieService) Create(req repository.Movie) (int, error) {
	movie := &repository.Movie{}

	err := movie.Create(&req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (MovieService) GetAll(page int, limit int, keyword string) ([]*repository.Movie, int, error) {
	movies := &repository.Movies{}

	result, err := movies.GetAll(page -1, limit, keyword)

	if err != nil {
		return *result, 500, err
	}

	return *result, 200, nil
}

func (MovieService) GetOne(id uint) (*repository.Movie, int, error) {
	movie := &repository.Movie{}

	result, err := movie.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (MovieService) Update(movie repository.Movie) {
	fmt.Println("Update a Movie")
}

func (MovieService) Delete(id uint) {
	fmt.Println("Delete a Movie")
}