package main

import (
	"log"

	"github.com/serge64/invite/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
