package main

import (
	"log"

	"github.com/abdullahnettoor/grpc-user-service/pkg/config"
	"github.com/abdullahnettoor/grpc-user-service/pkg/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("config error:", err.Error())
	}

	svr := server.NewServer()
	log.Println("Server Started on Port", cfg.Port)
	svr.Start(cfg)
}
