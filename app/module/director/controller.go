package director

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/go-http-std/app/model"
	"github.com/mrspec7er/go-http-std/app/utils"
)

type DirectorController struct {
	service DirectorService
}

func (c *DirectorController) HandlerCreate(w http.ResponseWriter, r *http.Request) {
	var director model.Director
	if err := json.NewDecoder(r.Body).Decode(&director); err != nil {
		utils.BadRequestHandler(w)
		return
	}

	status, err := c.service.Create(&director)
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.MutationSuccessResponse(w, "Successfully insert new Director")
}

func (c *DirectorController) HandlerGetAll(w http.ResponseWriter, r *http.Request) {
	result, status, err := c.service.GetAll()
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.GetSuccessResponse(w, nil, result, nil)
}

func (c *DirectorController) HandlerGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
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
