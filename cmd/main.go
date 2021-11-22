package main

import (
	"flag"
	"log"

	"github.com/ZeroXKiritsu/urlshortener/internal"
	"github.com/ZeroXKiritsu/urlshortener/internal/handlers"
	"github.com/ZeroXKiritsu/urlshortener/internal/repository"
	"github.com/ZeroXKiritsu/urlshortener/internal/service"
)

func main() {
	var databaseType string
	flag.StringVar(&databaseType, "db", "postgres", "database type")
	flag.Parse()

	db := repository.NewDatabase(databaseType)

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handlers.NewHandler(services)

	server := new(internal.Server)
	if err := server.Run(handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}