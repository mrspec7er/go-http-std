package main

import (
	"context"

	"github.com/mrspec7er/go-http-std/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		panic(err)
	} 
}