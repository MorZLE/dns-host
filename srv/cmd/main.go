package main

import (
	"dns-host/srv/config"
	"dns-host/srv/internal/grpc"
	"dns-host/srv/internal/service"
	"log/slog"
	"os"
	"strconv"
)

func main() {
	cnf := config.NewConfig()
	log := slog.Default()
	app := NewApp(log, cnf)

	app.MustRun()
}

func NewApp(log *slog.Logger, cfg *config.Config) *grpc.App {
	log.Info("starting server", slog.String("port", strconv.Itoa(cfg.GRPC.Port)))

	dns, err := service.NewDNSWorker(log, cfg.DNS.PathResolve)
	if err != nil {
		log.Error("failed to create dns", err)
		os.Exit(1)
	}
	logic := service.NewService(log, dns)
	grpcApp := grpc.NewGRPC(log, cfg.GRPC.Port, &logic)

	return grpcApp
}
