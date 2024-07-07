package service

import (
	"fmt"
	"math"

	"github.com/abdullahnettoor/grpc-user-service/pb"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
)

type userService struct {
	Users []*pb.User
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer() *userService {
	engine := grpc.NewServer()

	pb.RegisterUserServiceServer(engine, pb.UnimplementedUserServiceServer{})

	UserList := SimulateUserList()

	return &userService{
		Users: UserList,
	}
}

func SimulateUserList() []*pb.User {
	f := gofakeit.New(0)

	userList := make([]*pb.User, 100)

	fmt.Println("---\t\tUSER LIST\t\t---")
	fmt.Printf("ID  NAME \t CITY \t\t\t PHONE \t\t HEIGHT    MARRIED \n")

	for i := range userList {
		userList[i] = &pb.User{
			Id:        uint64(i + 1),
			Fname:     f.FirstName(),
			City:      f.City(),
			Phone:     f.Phone(),
			Height:    math.Round(f.Float64Range(4.5, 6.5)*100) / 100,
			IsMarried: f.Bool(),
		}
		fmt.Printf("%d   %s \t %s \t\t %s \t %.1f    %t \n", i+1, userList[i].Fname, userList[i].City, userList[i].Phone, userList[i].Height, userList[i].IsMarried)

	}

	return userList
}
