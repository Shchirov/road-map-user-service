package userservice

import (
	"context"
	"errors"
	"fmt"
	userv1 "github.com/Shchirov/road-map-api/gen/go/user"
	"log/slog"
	"road-map-user-server/internal/handler/grpc"
	"road-map-user-server/internal/model"
	userrepository "road-map-user-server/internal/repository/user"
	"road-map-user-server/internal/storage"
)

var _ usergrpc.UserService = Service{}

type Service struct {
	log            *slog.Logger
	userRepository *userrepository.Repository
}

func (u Service) CreateUser(ctx context.Context, request *userv1.CreateUserRequest) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) GetUsersById(ctx context.Context, userId string) (*model.User, error) {

	const op = "User.GetUsersById"
	user, err := u.userRepository.GetUserById(userId)

	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			u.log.With("op", op).Warn("user not found")
			return nil, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
	}

	return user, nil
}

func NewService(log *slog.Logger, repository *userrepository.Repository) *Service {
	return &Service{
		log:            log,
		userRepository: repository}
}
