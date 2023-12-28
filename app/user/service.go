package user

import (
	"errors"
	"fmt"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type UserService struct {}

func (UserService) Create(req repository.User) (int, error) {
	user := &repository.User{}

	err := user.Create(&req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (UserService) GetAll() ([]*repository.User, int, error) {
	users := &repository.Users{}

	result, err := users.GetAll()

	if err != nil {
		return *result, 500, err
	}

	return *result, 200, nil
}

func (UserService) GetOne(id uint) (*repository.User, int, error) {
	user := &repository.User{}

	result, err := user.GetByID(id)
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