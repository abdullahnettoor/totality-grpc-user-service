package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abdullahnettoor/grpc-user-service/pb"
	"github.com/abdullahnettoor/grpc-user-service/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Load configuration from a file or environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Error Ocurred in LoadConfig():", err.Error())
	}

	// Create a new gRPC client connection to the server
	conn, err := grpc.NewClient("localhost"+cfg.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error Ocurred in gRpcNewClient():", err.Error())
	}
	defer conn.Close()

	// Create a new UserService client from the connection
	client := pb.NewUserServiceClient(conn)
	log.Println("Client Created")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call the GetUserById method on the server
	fmt.Println("Get User By Id 5")
	userRes, err := client.GetUserByID(ctx, &pb.GetUserByIDReq{Id: 5})
	if err != nil {
		log.Printf("Error Ocurred: %v", err)
	}
	fmt.Println("Got Response: ", userRes)


	// Call the GetUserListByIds method on the server
	fmt.Println("Get Users By with ID 25, 3, 5, 99")
	userListRes, err := client.GetUserListByIDs(ctx, &pb.GetUserListByIDsReq{Ids: []uint64{25, 3, 5, 99}})
	if err != nil {
		log.Printf("Error Ocurred: %v", err)
	}
	fmt.Println("Got Response: ", userListRes)

	// Call the SearchUser method on the server #1
	fmt.Println("Search Users with Fname Contains with 'ab'")
	userSearchRes, err := client.SearchUser(ctx, &pb.SearchUserReq{Fname: "ab"})
	if err != nil {
		log.Printf("Error Ocurred: %v", err)
	}
	fmt.Println("Got Response: ", userSearchRes)
	
	// Call the SearchUser method on the server #1
	fmt.Println("Search Married Users")
	userSearchRes, err = client.SearchUser(ctx, &pb.SearchUserReq{IsMarriedFilterApplied: true, IsMarried: true})
	if err != nil {
		log.Printf("Error Ocurred: %v", err)
	}
	fmt.Println("Got Response: ", userSearchRes)
}
