package cmd

import (
	"Jinder/jinder-api/pkg/handler"
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service"
	"log"
)

func main() {
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal("Database connection error")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes()
}
