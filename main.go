package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mrspec7er/go-http-std/app"
)

func main() {
	app := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		panic(err)
	} 
}