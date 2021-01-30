package main

import (
	"Jinder/jinder-api/recommendations/pkg/consumer"
	"Jinder/jinder-api/recommendations/pkg/infrastructure"
	"Jinder/jinder-api/recommendations/pkg/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//nats://localhost:4222

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	infr := infrastructure.NewInfrastructure(viper.GetString("clients.jobs_service"))
	serv := service.NewService(infr)
	cons := consumer.NewConsumer(serv)

	cons.InitRoutes(viper.GetString("clients.nats"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
