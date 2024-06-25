package main

import (
	"dns-host/srv/config"
	"dns-host/srv/internal"
	"dns-host/srv/internal/service"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	cnf := config.NewConfig()
	log := slog.Default()
	app := NewApp(log, cnf)

	go app.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Info("stopping application", slog.String("signal", sig.String()))

	app.Stop()
	log.Info("application stop")
}

func NewApp(log *slog.Logger, cfg *config.Config) *internal.App {
	log.Info("starting server", slog.String("port", strconv.Itoa(cfg.GRPC.Port)))

	dns, err := service.NewDNSWorker(log)
	if err != nil {
		log.Error("failed to create dns", err)
		os.Exit(1)
	}
	logic := service.NewService(log, dns)
	grpcApp := internal.NewGRPC(log, cfg.GRPC.Port, &logic)

	return grpcApp
}
