package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"road-map-user-server/internal/app"
	"road-map-user-server/internal/config"
	"road-map-user-server/internal/storage/pg"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("reading config")
	cfg := config.MustLoad()

	log.Println("initialize logger")
	lg := SetupLogger(cfg.Env)

	ctx := context.Background()

	db := pg.Dial(&cfg.AppConfig.DataBaseConfig)

	grpcServerApp := app.NewGrpcServer(lg, db, &cfg.AppConfig.GrpcConfig)

	httpServerApp := app.NewHttpServer(lg, ctx, cfg.AppConfig.HttpConfig, cfg.AppConfig.GrpcConfig)

	go func() {
		grpcServerApp.GrpcServer.MustRun()
	}()

	go func() {
		httpServerApp.HttpServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	grpcServerApp.GrpcServer.Stop()
	httpServerApp.HttpServer.Stop()
	lg.Info("Gracefully stopped")
	return nil
}

const (
	envLocal = "local"
	envProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
