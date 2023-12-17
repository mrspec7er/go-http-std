package main

import (
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(handler),
	}

	log.Println("Server Start On Port :", 8080)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}

}

func handler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello There!"))
}