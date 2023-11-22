package main

import (
	"log"

	gostudio "github.com/camarin24/go-studio"
)

func main() {
	app := gostudio.NewWithConfig()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
