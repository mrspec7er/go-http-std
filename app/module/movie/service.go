package movie

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/mrspec7er/go-http-std/app/repository"
	"gorm.io/gorm"
)

type MovieService struct {
	movie repository.Movie
}

func (s *MovieService) Create(req *repository.Movie) (int, error) {
	s.movie = *req
	err := s.movie.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *MovieService) GetAll(page int, limit int, keyword string) ([]*repository.Movie, *int64, int, error) {

	result, count, err := s.movie.GetAll(page -1, limit, keyword)

	if err != nil {
		return result, nil, 500, err
	}

	return result, count, 200, nil
}

func (s *MovieService) GetOne(id uint) (*repository.Movie, int, error) {
	result, err := s.movie.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, 400, err
	}

	if err != nil {
		return result, 500, err
	}

	return result, 200, nil
}

func (s *MovieService) Update(req *repository.Movie) (int, error) {
	s.movie = *req
	err := s.movie.Update()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 400, err
	}

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *MovieService) UpdateThumbnail(file multipart.File, fileHeader *multipart.FileHeader, id uint) error {
	s.RemoveMovieThumbnail(id)

	err := os.MkdirAll("./assets/thumbnail", os.ModePerm)
	if err != nil {
		return err
	}

	dst, err := os.Create(fmt.Sprintf("./assets/thumbnail/%d_%s", time.Now().UnixNano(), fileHeader.Filename))
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	thumbnailUrl := strings.TrimPrefix(dst.Name(), ".")
	s.movie.ID = id
	s.movie.Thumbnail = thumbnailUrl

	s.movie.Update()
	return nil
}

func (s *MovieService) RemoveMovieThumbnail(id uint) error {
	movie, err := s.movie.GetByID(id)
	if err != nil {
		return err
	}

	dst := "." + movie.Thumbnail
	err = os.Remove(dst)
	if err != nil {
		return err
	}

	return nil
}

func (MovieService) Delete(id uint) {
	fmt.Println("Delete a Movie")
}