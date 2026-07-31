package handler

import (
	"context"

	"grpc/internal/model"
	"grpc/internal/service"
	pb "grpc/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	user := &model.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	err := h.service.Create(user)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:    uint32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, req *pb.Empty) (*pb.UserListResponse, error) {

	users, err := h.service.GetAll()
	if err != nil {
		return nil, err
	}

	var userList []*pb.User

	for _, user := range users {
		userList = append(userList, &pb.User{
			Id:    uint32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &pb.UserListResponse{
		Users: userList,
	}, nil
}