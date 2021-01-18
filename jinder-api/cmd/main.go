package main

import (
	"Jinder/jinder-api"
	"Jinder/jinder-api/pkg/handler"
	"Jinder/jinder-api/pkg/repository"
	"Jinder/jinder-api/pkg/service"
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}
	password, ok := os.LookupEnv("DB_PASSWORD")

	if !ok {
		log.Fatal("Error while getting db password environment")
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Username: viper.GetString("db.username"),
		Password: password,
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatal("Database connection error", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(jinder.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("jinder-api/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
