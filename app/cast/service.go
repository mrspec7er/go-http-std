package cast

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type CastService struct {}

func (CastService) Create(req repository.Cast) (int, error) {
	cast := &repository.Cast{}

	err := cast.Create(&req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (CastService) GetAll() ([]*repository.Cast, int, error) {
	casts := &repository.Casts{}

	result, err := casts.GetAll()

	if err != nil {
		return *result, 500, err
	}

	return *result, 200, nil
}

func (CastService) GetOne(id uint) (*repository.Cast, int, error) {
	cast := &repository.Cast{}

	result, err := cast.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (CastService) Update(cast repository.Cast) {
	fmt.Println("Update a cast")
}

func (CastService) Delete(id uint) {
	fmt.Println("Delete a cast")
}