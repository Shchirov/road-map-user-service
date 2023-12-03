package usergrpc

import (
	"context"
	"errors"
	"fmt"
	userv1 "github.com/Shchirov/road-map-api/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"road-map-user-server/internal/model"
	"road-map-user-server/internal/storage"
)

type UserService interface {
	CreateUser(ctx context.Context, request *userv1.CreateUserRequest) (model.User, error)
	GetUsersById(ctx context.Context, userId string) (*model.User, error)
}

type UserServer struct {
	userv1.UnimplementedUserServer
	userService UserService
}

func (u UserServer) CreateUser(ctx context.Context, request *userv1.CreateUserRequest) (*userv1.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) GetUsersById(ctx context.Context, request *userv1.GetUserByIdRequest) (*userv1.UserResponse, error) {

	if request.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	user, err := u.userService.GetUsersById(ctx, request.Id)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("user with id  = %s not found ", request.GetId()))
		}
	}

	return user.ToResponse(), nil
}

func (u UserServer) UpdateUser(ctx context.Context, request *userv1.UpdateUserRequest) (*userv1.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func Register(gRPCServer *grpc.Server, us UserService) {
	userv1.RegisterUserServer(gRPCServer, &UserServer{userService: us})
}
