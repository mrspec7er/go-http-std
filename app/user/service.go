package user

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type UserService struct {
	user repository.User
}

func (s *UserService) Create(req *repository.User) (int, error) {
	s.user = repository.User{Name: req.Name, Email: req.Email, Password: "UNFILLED", Status: "INACTIVE", Role: "USER"}

	err := s.user.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *UserService) GetAll() ([]*repository.User, int, error) {
	result, err := s.user.GetAll()

	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (s *UserService) GetOne(id uint) (*repository.User, int, error) {
	result, err := s.user.GetByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}
	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (UserService) Update(User repository.User) {
	fmt.Println("Update a User")
}

func (UserService) Delete(id uint) {
	fmt.Println("Delete a User")
}