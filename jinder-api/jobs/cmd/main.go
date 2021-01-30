package main

import (
	"Jinder/jinder-api/jobs/pkg/handler/http"
	"Jinder/jinder-api/jobs/pkg/infrastructure"
	"Jinder/jinder-api/jobs/pkg/repository"
	"Jinder/jinder-api/jobs/pkg/service"
	"context"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}
	password, ok := os.LookupEnv("DB_PASSWORD")

	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	if !ok {
		log.Fatal("Error while getting db password environment")
	}
	config := repository.Config{
		Username: viper.GetString("db.username"),
		Password: password,
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	infrastructureConfig := infrastructure.Config{
		NatsUrl: viper.GetString("clients.nats"),
	}

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	nats, err := infrastructure.NewNatsConnection(infrastructureConfig.NatsUrl)
	if err != nil {
		log.Fatal("Nats connection error", err)
	}

	repos := repository.NewRepository(db)
	infr := infrastructure.NewInfrastructure(nats)
	services := service.NewService(repos, infr)
	handlers := http.NewHandler(services)
	//grpc := grpc.NewJobsServer(services)

	srv := new(http.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print("Jinder Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Jinder Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
