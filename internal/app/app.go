package app

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	grpcapp "road-map-user-server/internal/app/grpc"
	httpapp "road-map-user-server/internal/app/http"
	"road-map-user-server/internal/config"
	userrepository "road-map-user-server/internal/repository/user"
	"road-map-user-server/internal/service/user"
)

type GrpcApp struct {
	GrpcServer *grpcapp.App
}

type HttpApp struct {
	HttpServer *httpapp.App
}

func NewGrpcServer(log *slog.Logger, db *gorm.DB, grpcConfig *config.GrpcConfig) *GrpcApp {

	repository := userrepository.NewRepository(db, log)
	userService := userservice.NewService(log, repository)

	grpcApp := grpcapp.New(log, userService, grpcConfig)

	return &GrpcApp{
		GrpcServer: grpcApp,
	}
}

func NewHttpServer(
	log *slog.Logger,
	ctx context.Context,
	httpConfig config.HttpConfig,
	grpcClientConfig config.GrpcConfig,
) *HttpApp {

	httpApp := httpapp.New(log, ctx, httpConfig, grpcClientConfig)

	return &HttpApp{
		HttpServer: httpApp,
	}
}
