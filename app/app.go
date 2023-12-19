package app

import (
	"context"
	"net/http"

	"github.com/mrspec7er/go-http-std/app/model"
	"github.com/mrspec7er/go-http-std/app/utils"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{
		router: loadRoutes(),
	}
}


func (a *App) Start(ctx context.Context) error {

	utils.DBConnection()
	
	utils.DB.AutoMigrate(&model.Movie{})
	
	server := &http.Server{
		Addr: ":8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	
	return nil
}

