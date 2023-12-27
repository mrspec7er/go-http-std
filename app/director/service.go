package director

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type DirectorService struct {}

func (DirectorService) Create(req repository.Director) (int, error) {
	director := &repository.Director{}

	err := director.Create(&req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (DirectorService) GetAll() ([]*repository.Director, int, error) {
	directors := &repository.Directors{}

	result, err := directors.GetAll()

	if err != nil {
		return *result, 500, err
	}

	return *result, 200, nil
}

func (DirectorService) GetOne(id uint) (*repository.Director, int, error) {
	director := &repository.Director{}

	result, err := director.GetByID(id)
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