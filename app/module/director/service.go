package director

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type DirectorService struct {
	director repository.Director
}

func (s *DirectorService) Create(req *repository.Director) (int, error) {
	s.director = *req

	err := s.director.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *DirectorService) GetAll() ([]*repository.Director, int, error) {
	result, err := s.director.GetAll()

	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (s *DirectorService) GetOne(id uint) (*repository.Director, int, error) {
	result, err := s.director.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (DirectorService) Update(director repository.Director) {
	fmt.Println("Update a director")
}

func (DirectorService) Delete(id uint) {
	fmt.Println("Delete a director")
}