package user

import (
	"context"
	"errors"
)

// In-memory user storage
var users = []User{
	{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	// Add more sample users as needed
}

// UserServer implementation
type UserServer struct {
	UnimplementedUserServiceServer
}

func (s *UserServer) GetUserByID(ctx context.Context, req *UserIDRequest) (*UserResponse, error) {
	for _, user := range users {
		if user.Id == req.Id {
			return &UserResponse{User: &user}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserServer) GetUsersByID(ctx context.Context, req *UserIDsRequest) (*UsersResponse, error) {
	var foundUsers []*User
	for _, id := range req.Ids {
		for _, user := range users {
			if user.Id == id {
				foundUsers = append(foundUsers, &user)
			}
		}
	}
	return &UsersResponse{Users: foundUsers}, nil
}

func (s *UserServer) SearchUsers(ctx context.Context, req *User) (*UsersResponse, error) {
	var foundUsers []*User
	for _, user := range users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(req.Married == user.Married) {
			foundUsers = append(foundUsers, &user)
		}
	}
	return &UsersResponse{Users: foundUsers}, nil
}
