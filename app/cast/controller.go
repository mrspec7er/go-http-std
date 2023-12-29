package cast

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/go-http-std/app/repository"
	"github.com/mrspec7er/go-http-std/app/utils"
)


type CastController struct {
	service CastService
}

func(c *CastController) HandlerCreate(w http.ResponseWriter, r *http.Request)  {
	var cast repository.Cast
    if err := json.NewDecoder(r.Body).Decode(&cast); err != nil {
        utils.BadRequestHandler(w)
        return
    }

	status, err := c.service.Create(&cast)
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.MutationSuccessResponse(w, "Successfully insert new Cast")
}

func (c *CastController) HandlerGetAll(w http.ResponseWriter, r *http.Request) {
	result, status, err := c.service.GetAll()
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.GetSuccessResponse(w, nil, result, nil)
}

func (c *CastController) HandlerGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id");
	formattedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.BadRequestHandler(w)
        return
	}

	result, status, err := c.service.GetOne(uint(formattedId))
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.GetSuccessResponse(w, nil, result, nil)
}

func HandlerUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a Movie")
}

func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a Movie")
}