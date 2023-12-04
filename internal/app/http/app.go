package httpapp

import (
	"context"
	"fmt"
	userv1 "github.com/Shchirov/road-map-api/gen/go/user"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
	"net/http"
	"road-map-user-server/internal/config"
)

type App struct {
	log          *slog.Logger
	mux          *runtime.ServeMux
	httpListener net.Listener
	httpConfig   config.HttpConfig
}

// New creates new Http server app.
func New(
	log *slog.Logger,
	ctx context.Context,
	httpConfig config.HttpConfig,
	grpcClientConfig config.GrpcConfig,
) *App {

	mux := runtime.NewServeMux()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", grpcClientConfig.GrpcHost, grpcClientConfig.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("todo")
	}

	httpListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", httpConfig.HttpHost, httpConfig.HttpPort))
	if err != nil {
		panic("todo")
	}

	err = userv1.RegisterUserHandler(ctx, mux, conn)
	if err != nil {
		panic("todo")
	}

	return &App{
		log:          log,
		httpListener: httpListener,
		httpConfig:   httpConfig,
		mux:          mux,
	}
}

// MustRun runs http server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs http server.
func (a *App) Run() error {
	const op = "htppapp.Run"

	a.log.With(slog.String("op", op)).Info("http server started", slog.String("addr", fmt.Sprintf("%s:%d", a.httpConfig.HttpHost, a.httpConfig.HttpPort)))

	if err := http.Serve(a.httpListener, a.mux); err != nil {
		panic("todo")
	}

	return nil
}

// Stop stops http server.
func (a *App) Stop() {
	const op = "httpapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping http server", slog.Int("port", a.httpConfig.HttpPort))

	err := a.httpListener.Close()
	if err != nil {
		return
	}
}
