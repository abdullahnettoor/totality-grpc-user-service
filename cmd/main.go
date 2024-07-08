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
	// Load the configuration from a env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("config error:", err.Error())
	}

    // Create a TCP listener on the port specified in the config
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("error ocurred:", err.Error())
	}

    // Create and Register the UserService with the server
	svr := grpc.NewServer()
	pb.RegisterUserServiceServer(svr, service.NewUserServiceServer())

    // Register the reflection service to allow for service discovery
	reflection.Register(svr)

	log.Println("Server Started on Port", cfg.Port)
	// Start serving requests on the listener
	if err := svr.Serve(lis); err != nil {
		log.Fatalln("server error ocurred:", err.Error())
	}
}
