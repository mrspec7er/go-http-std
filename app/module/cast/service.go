package cast

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/model"
	"gorm.io/gorm"
)

type CastService struct {
	cast model.Cast
}

func (s *CastService) Create(req *model.Cast) (int, error) {
	s.cast = *req

	err := s.cast.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *CastService) GetAll() ([]*model.Cast, int, error) {
	result, err := s.cast.GetAll()

	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (s *CastService) GetOne(id uint) (*model.Cast, int, error) {

	result, err := s.cast.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (CastService) Update(cast model.Cast) {
	fmt.Println("Update a cast")
}

func (CastService) Delete(id uint) {
	fmt.Println("Delete a cast")
}
