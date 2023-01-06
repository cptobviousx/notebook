package main

import (
	"log"

	"github.com/cptobviousx/notebook"
	"github.com/cptobviousx/notebook/pkg/handler"
	"github.com/cptobviousx/notebook/pkg/repository"
	"github.com/cptobviousx/notebook/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(notebook.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
