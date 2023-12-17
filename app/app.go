package app

import (
	"context"
	"net/http"
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