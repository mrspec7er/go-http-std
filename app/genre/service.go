package genre

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type GenreService struct {
	genre repository.Genre
}

func (s *GenreService) Create(req *repository.Genre) (int, error) {
	s.genre = *req
	err := s.genre.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *GenreService) GetAll() ([]*repository.Genre, int, error) {

	result, err := s.genre.GetAll()

	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (s *GenreService) GetOne(id uint) (*repository.Genre, int, error) {
	s.genre.ID = id

	result, err := s.genre.GetByID(id)
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