package main

import (
	"log"

	"market/internal/pkg/application"
)

func main() {
	app, err := application.New()
	if err != nil {
		log.Printf("%+v\n", err)
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}
}
