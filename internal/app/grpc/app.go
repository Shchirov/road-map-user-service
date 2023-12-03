package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"road-map-user-server/internal/config"
	"road-map-user-server/internal/handler/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	grpcConfig *config.GrpcConfig
}

// New creates new gRPC server app.
func New(
	log *slog.Logger,
	userService usergrpc.UserService,
	grpcConfig *config.GrpcConfig,
) *App {

	gRPCServer := grpc.NewServer()

	usergrpc.Register(gRPCServer, userService)

	return &App{
		grpcConfig: grpcConfig,
		log:        log,
		gRPCServer: gRPCServer,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcConfig.GrpcPort))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.
		With(slog.String("op", op)).
		Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", a.grpcConfig.GrpcPort))

	a.gRPCServer.GracefulStop()
}
