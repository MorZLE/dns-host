package grpc

import (
	grpcServer "dns-host/gen/server"
	"dns-host/srv/internal/service"
	"google.golang.org/grpc"
	"log/slog"
)

func NewController(log *slog.Logger, service *service.IService) *ServerAPI {
	return &ServerAPI{
		srv: *service,
		log: log,
	}
}

type ServerAPI struct {
	grpcServer.UnimplementedServiceDNSServer
	srv service.IService
	log *slog.Logger
}

func RegisterServerAPI(gRPC *grpc.Server, srv *ServerAPI) {
	grpcServer.RegisterServiceDNSServer(gRPC, srv)
}
