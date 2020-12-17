package main

import (
	"github.com/Thunderbirrd/Go_TODO"
	"github.com/Thunderbirrd/Go_TODO/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil{
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
