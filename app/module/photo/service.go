package photo

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/mrspec7er/go-http-std/app/model"
)

type PhotoService struct {
	photo model.Photo
}

func (s *PhotoService) Creates(req []*model.Photo) (int, error) {

	err := s.photo.BulkCreate(req)
	if err != nil {
		return 500, err
	}

	return 200, nil
}

func (s *PhotoService) UploadPhoto(files []*multipart.FileHeader, movieId uint) ([]*model.Photo, error) {
	photos := []*model.Photo{}

	for _, fileHeader := range files {
		err := os.MkdirAll("./assets/photos", os.ModePerm)
		if err != nil {
			return nil, err
		}

		dst, err := os.Create(fmt.Sprintf("./assets/photos/%d_%s", time.Now().UnixNano(), fileHeader.Filename))
		if err != nil {
			return nil, err
		}

		defer dst.Close()

		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(dst, file)
		if err != nil {
			return nil, err
		}

		photoUrl := strings.TrimPrefix(dst.Name(), ".")
		photos = append(photos, &model.Photo{Name: fileHeader.Filename, URL: photoUrl, MovieID: movieId})
	}

	return photos, nil
}

func (s *PhotoService) GetByMovie(movieId uint) ([]*model.Photo, error) {
	result, err := s.photo.GetByMovieId(movieId)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *PhotoService) DeleteByMovie(movieId uint) (int, error) {
	photos, err := s.photo.GetByMovieId(movieId)
	if err != nil {
		return 400, err
	}

	fmt.Println(movieId, photos)

	for _, p := range photos {
		dst := "." + p.URL
		err = os.Remove(dst)
		if err != nil {
			return 500, err
		}
	}

	err = s.photo.DeleteByMovie(movieId)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
