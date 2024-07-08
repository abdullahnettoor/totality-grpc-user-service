package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/abdullahnettoor/grpc-user-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_userService_GetUserByID(t *testing.T) {
	type fields struct {
		Users                          []*pb.User
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.GetUserByIDReq
	}
	userList := SimulateUserList(100)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetUserByIDRes
		wantErr bool
	}{
		{
			name:    "Valid User ID",
			fields:  fields{Users: userList},
			args:    args{context.Background(), &pb.GetUserByIDReq{Id: 100}},
			want:    &pb.GetUserByIDRes{User: userList[99]},
			wantErr: false,
		},
		{
			name:    "User ID Don't Exist",
			fields:  fields{Users: userList},
			args:    args{context.Background(), &pb.GetUserByIDReq{Id: 0}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Empty User List",
			fields:  fields{Users: SimulateUserList(0)},
			args:    args{context.Background(), &pb.GetUserByIDReq{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid Request",
			fields:  fields{Users: userList},
			args:    args{context.Background(), &pb.GetUserByIDReq{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				Users:                          tt.fields.Users,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := us.GetUserByID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserListByIDs(t *testing.T) {
	type fields struct {
		Users                          []*pb.User
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.GetUserListByIDsReq
	}
	userList := SimulateUserList(100)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetUserListByIDsRes
		wantErr error
	}{
		{
			name:   "Users with Valid IDs",
			fields: fields{Users: userList},
			args:   args{context.Background(), &pb.GetUserListByIDsReq{Ids: []uint64{1, 4, 66, 78, 99}}},
			want: &pb.GetUserListByIDsRes{UserList: []*pb.User{
				userList[0],
				userList[3],
				userList[65],
				userList[77],
				userList[98]}},
			wantErr: nil,
		},
		{
			name:   "User with Valid and Invalid IDs",
			fields: fields{Users: userList},
			args:   args{context.Background(), &pb.GetUserListByIDsReq{Ids: []uint64{0, 8, 19, 46, 81, 101, 1000, 898}}},
			want: &pb.GetUserListByIDsRes{UserList: []*pb.User{
				userList[7],
				userList[18],
				userList[45],
				userList[80]},
				NotFound: []uint64{0, 101, 1000, 898}},
			wantErr: nil,
		},
		{
			name:    "Users Not Found with IDs",
			fields:  fields{Users: userList},
			args:    args{context.Background(), &pb.GetUserListByIDsReq{Ids: []uint64{0, 101, 201, 150}}},
			want:    nil,
			wantErr: status.Errorf(codes.NotFound, "users not found"),
		},
		{
			name:    "Empty Request",
			fields:  fields{Users: userList},
			args:    args{context.Background(), &pb.GetUserListByIDsReq{}},
			want:    nil,
			wantErr: status.Errorf(codes.InvalidArgument, "user ids not provided"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				Users:                          tt.fields.Users,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := us.GetUserListByIDs(tt.args.ctx, tt.args.req)
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("\nuserService.GetUserListByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nuserService.GetUserListByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_SearchUser(t *testing.T) {
	type fields struct {
		Users                          []*pb.User
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		ctx context.Context
		req *pb.SearchUserReq
	}
	userList := SimulateUserList(100)
	marriedList := []*pb.User{}
	for _, u := range userList {
		if u.IsMarried {
			marriedList = append(marriedList, u)
		}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SearchUserRes
		wantErr error
	}{
		{
			name:   "Filter Users",
			fields: fields{Users: userList},
			args: args{context.Background(), &pb.SearchUserReq{
				Fname: userList[0].Fname[2:],
				City:  userList[0].City[:3],
			}},
			want: &pb.SearchUserRes{UserList: []*pb.User{
				userList[0],
			}},
			wantErr: nil,
		},
		{
			name:   "Filter By Phone",
			fields: fields{Users: userList},
			args: args{context.Background(), &pb.SearchUserReq{
				Phone: userList[9].Phone,
			}},
			want:    &pb.SearchUserRes{UserList: []*pb.User{userList[9]}},
			wantErr: nil,
		},
		{
			name:   "Filter By Married",
			fields: fields{Users: userList},
			args: args{context.Background(), &pb.SearchUserReq{
				IsMarriedFilterApplied: true,
				IsMarried:              true,
			}},
			want:    &pb.SearchUserRes{UserList: marriedList},
			wantErr: nil,
		},
		{
			name:   "Not Found Filter Users",
			fields: fields{Users: userList},
			args: args{context.Background(), &pb.SearchUserReq{
				Fname: userList[0].Fname[2:],
				City:  userList[1].City[:3],
				Phone: userList[5].Phone,
			}},
			want:    nil,
			wantErr: status.Errorf(codes.NotFound, "user not found with applied filter"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				Users:                          tt.fields.Users,
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := us.SearchUser(tt.args.ctx, tt.args.req)
			if err != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("\nuserService.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nuserService.SearchUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
