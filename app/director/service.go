package director

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type DirectorService struct {}

func (DirectorService) Create(req repository.Director) (int, error) {
	movie := &repository.Director{}

	err := movie.Create(&req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (DirectorService) GetAll() ([]*repository.Director, int, error) {
	movies := &repository.Directors{}

	result, err := movies.GetAll()

	if err != nil {
		return *result, 500, err
	}

	return *result, 200, nil
}

func (DirectorService) GetOne(id uint) (*repository.Director, int, error) {
	movie := &repository.Director{}

	result, err := movie.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (DirectorService) Update(movie repository.Director) {
	fmt.Println("Update a Movie")
}

func (DirectorService) Delete(id uint) {
	fmt.Println("Delete a Movie")
}