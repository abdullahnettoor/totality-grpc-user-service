package service

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/abdullahnettoor/grpc-user-service/pb"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	Users []*pb.User
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer() *userService {
	engine := grpc.NewServer()

	pb.RegisterUserServiceServer(engine, pb.UnimplementedUserServiceServer{})

	UserList := SimulateUserList(100)

	return &userService{
		Users: UserList,
	}
}

func SimulateUserList(count uint) []*pb.User {
	f := gofakeit.New(0)

	userList := make([]*pb.User, count)

	fmt.Println("--------------------------------USER LIST--------------------------------")
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

func (us *userService) GetUserByID(ctx context.Context, req *pb.GetUserByIDReq) (*pb.GetUserByIDRes, error) {

	res := &pb.GetUserByIDRes{}
	for _, user := range us.Users {
		if user.Id == req.Id {
			res.User = user
			return res, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "user not found")
}

func (us *userService) GetUserListByIDs(ctx context.Context, req *pb.GetUserListByIDsReq) (*pb.GetUserListByIDsRes, error) {

	if len(req.Ids) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "user ids not provided")
	}

	res := &pb.GetUserListByIDsRes{}

	notFound := []uint64{}
	for _, id := range req.Ids {
		found := false
		for _, user := range us.Users {
			if user.Id == id {
				res.UserList = append(res.UserList, user)
				found = true
				break
			}
		}
		if !found {
			notFound = append(notFound, id)
		}
	}

	if len(res.UserList) == 0 {
		return nil, status.Errorf(codes.NotFound, "users not found")
	} 

	fmt.Printf("Found %d results\n", len(res.UserList))

	if len(notFound) != 0 {
		res.NotFound = notFound
	}

	return res, nil
}

func (us *userService) SearchUser(ctx context.Context, req *pb.SearchUserReq) (*pb.SearchUserRes, error) {

	res := &pb.SearchUserRes{}

	for _, user := range us.Users {
		fnameMatch := req.Fname == "" || strings.Contains(strings.ToLower(user.Fname), strings.ToLower(req.Fname))
		cityMatch := req.City == "" || strings.Contains(strings.ToLower(user.City), strings.ToLower(req.City))
		phoneMatch := req.Phone == "" || strings.Contains(user.Phone, req.Phone)
		isMarriedMatch := !req.IsMarriedFilterApplied || user.IsMarried == req.IsMarried

		if fnameMatch && cityMatch && phoneMatch && isMarriedMatch {
			res.UserList = append(res.UserList, user)
		}
	}

	if len(res.UserList) == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found with applied filter")
	} else {
		fmt.Printf("Found %d results\n", len(res.UserList))
	}

	return res, nil
}
