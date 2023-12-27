package genre

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type GenreService struct {}

func (GenreService) Create(req repository.Genre) (int, error) {
	genre := &repository.Genre{}

	err := genre.Create(&req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (GenreService) GetAll() ([]*repository.Genre, int, error) {
	genres := &repository.Genres{}

	result, err := genres.GetAll()

	if err != nil {
		return *result, 500, err
	}

	return *result, 200, nil
}

func (GenreService) GetOne(id uint) (*repository.Genre, int, error) {
	genre := &repository.Genre{}

	result, err := genre.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (GenreService) Update(genre repository.Genre) {
	fmt.Println("Update a genre")
}

func (GenreService) Delete(id uint) {
	fmt.Println("Delete a genre")
}