package photo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/go-http-std/app/utils"
)

type PhotoController struct {
	service PhotoService
}

func(c *PhotoController) HandlerCreate(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieIdStringify := chi.URLParam(r, "movieId");

	movieId, err := strconv.ParseUint(movieIdStringify, 10, 32)
	if err != nil {
		utils.BadRequestHandler(w)
		return
	}
	
	fmt.Println(movieId)
	files := r.MultipartForm.File["photos"]

	photos, err := c.service.UploadPhoto(files, uint(movieId))

	status, err := c.service.Creates(photos)
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.MutationSuccessResponse(w, "Successfully insert new photos")
}

func (c *PhotoController) HandlerGetAll(w http.ResponseWriter, r *http.Request) {
	movieIdStringify := chi.URLParam(r, "movieId");

	movieId, err := strconv.ParseUint(movieIdStringify, 10, 32)
	if err != nil {
		utils.BadRequestHandler(w)
		return
	}

	result, err := c.service.GetByMovie(uint(movieId))
	if err != nil {
		utils.InternalServerErrorHandler(w, 500, err)
		return
	}

	utils.GetSuccessResponse(w, nil, result, nil)
}

func (c *PhotoController) HandlerDelete(w http.ResponseWriter, r *http.Request) {
	movieIdStringify := chi.URLParam(r, "movieId");

	movieId, err := strconv.ParseUint(movieIdStringify, 10, 32)
	if err != nil {
		utils.BadRequestHandler(w)
		return
	}

	status, err := c.service.DeleteByMovie(uint(movieId))
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.SuccessMessageResponse(w, "Successfully deleted movie photos")
}