package main

import (
	"github.com/Thunderbirrd/Go_TODO"
	"github.com/Thunderbirrd/Go_TODO/pkg/handler"
	"github.com/Thunderbirrd/Go_TODO/pkg/repository"
	"github.com/Thunderbirrd/Go_TODO/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil{
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers :=handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
