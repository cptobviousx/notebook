package main

import (
	"log"

	"github.com/cptobviousx/notebook"
	"github.com/cptobviousx/notebook/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(notebook.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
