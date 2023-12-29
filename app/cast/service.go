package cast

import (
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
)

type CastService struct {
	cast repository.Cast
}

func (s *CastService) Create(req *repository.Cast) (int, error) {
	s.cast = *req

	err := s.cast.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *CastService) GetAll() ([]*repository.Cast, int, error) {
	result, err := s.cast.GetAll()

	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (s *CastService) GetOne(id uint) (*repository.Cast, int, error) {

	result, err := s.cast.GetByID(id)
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