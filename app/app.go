package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/mrspec7er/go-http-std/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{
		router: loadRoutes(),
	}
}

func DBConnection() *gorm.DB {
	credentials := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"))

	connection, err := gorm.Open(postgres.Open(credentials), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return connection
}

func (a *App) Start(ctx context.Context) error {
	DB := DBConnection()
	DB.AutoMigrate(&model.Movie{})
	
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

