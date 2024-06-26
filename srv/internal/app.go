package internal

import (
	grpc2 "dns-host/srv/internal/grpc"
	"dns-host/srv/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func NewGRPC(log *slog.Logger, port int, service *service.IService) *App {
	grpcServer := grpc.NewServer()
	grpc2.RegisterServerAPI(grpcServer, grpc2.NewController(log, service))

	return &App{
		log:        log,
		port:       port,
		gRPCServer: grpcServer,
	}
}

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}

func (a *App) run() error {
	const op = "grpc.app.Run"
	log := a.log.With(slog.String("op", op))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("running controller server", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpc.app.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
