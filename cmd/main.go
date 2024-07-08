package main

import (
	"log"
	"net"

	"github.com/abdullahnettoor/grpc-user-service/pb"
	"github.com/abdullahnettoor/grpc-user-service/pkg/config"
	"github.com/abdullahnettoor/grpc-user-service/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("config error:", err.Error())
	}

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("error ocurred:", err.Error())
	}

	svr := grpc.NewServer()
	pb.RegisterUserServiceServer(svr, service.NewUserServiceServer())

	reflection.Register(svr)

	log.Println("Server Started on Port", cfg.Port)
	if err := svr.Serve(lis); err != nil {
		log.Fatalln("server error ocurred:", err.Error())
	}
}
