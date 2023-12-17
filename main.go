package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Println("Server Start On Port :", 8080)

	router.Get("/", handler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello There!"))
}