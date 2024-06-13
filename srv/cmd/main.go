package main

import (
	"dns-host/srv/config"
	"dns-host/srv/internal/grpc"
	"dns-host/srv/internal/service"
	"log/slog"
	"strconv"
)

func main() {
	cnf := config.NewConfig()
	log := slog.Default()
	app := NewApp(log, cnf)

	app.MustRun()
}

func NewApp(log *slog.Logger, cfg *config.Config) *grpc.App {
	log.Info("starting server", slog.String("port", strconv.Itoa(cfg.Port)))

	logic := service.NewService(log)

	grpcApp := grpc.NewGRPC(log, cfg.Port, &logic)

	return grpcApp
}
