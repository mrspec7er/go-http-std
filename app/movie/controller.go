package movie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/go-http-std/app/repository"
	"github.com/mrspec7er/go-http-std/app/utils"
)

type MovieController struct {
	Movie MovieService
}

func(c *MovieController) HandlerCreate(w http.ResponseWriter, r *http.Request)  {
	var movie repository.Movie
    if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
        utils.BadRequestHandler(w)
        return
    }

	status, err := c.Movie.Create(movie)
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	utils.MutationSuccessResponse(w, "Successfully insert new movie")

}

func (c *MovieController) HandlerGetAll(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		utils.BadRequestHandler(w)
        return
	}

	l := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(l)
	if err != nil {
		utils.BadRequestHandler(w)
        return
	}

	keyword := r.URL.Query().Get("keyword")
	
	result, count, status, err := c.Movie.GetAll(page, limit, keyword)
	if err != nil {
		utils.InternalServerErrorHandler(w, status, err)
		return
	}

	metadata := utils.Metadata{
		Page: page,
		Limit: limit,
		Count: *count,
	}

	utils.GetSuccessResponse(w, nil, result, &metadata)
}

func (c *MovieController) HandlerGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id");
	formattedId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.BadRequestHandler(w)
        return
	}

	result, status, err := c.Movie.GetOne(uint(formattedId))
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