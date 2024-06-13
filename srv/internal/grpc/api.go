package grpc

import (
	"context"
	grpcServer "dns-host/gen/server"
	"dns-host/srv/internal/service"

	"google.golang.org/grpc"
)

func NewController(service *service.IService) *ServerAPI {
	return &ServerAPI{
		srv: service,
	}

}

type ServerAPI struct {
	grpcServer.UnimplementedServiceDNSServer
	srv *service.IService
}

func RegisterServerAPI(gRPC *grpc.Server, srv *ServerAPI) {
	grpcServer.RegisterServiceDNSServer(gRPC, srv)
}

func (s *ServerAPI) SetHostname(context.Context, *grpcServer.SetHostnameRequest) (*grpcServer.SetHostnameResponse, error) {
	return nil, nil
}
func (s *ServerAPI) GetHostname(context.Context, *grpcServer.GetHostnameRequest) (*grpcServer.GetHostnameResponse, error) {
	return nil, nil
}
func (s *ServerAPI) GetAllDNS(context.Context, *grpcServer.GetAllDNSRequest) (*grpcServer.GetAllDNSResponse, error) {
	return nil, nil
}
func (s *ServerAPI) SetDNS(context.Context, *grpcServer.SetDNSRequest) (*grpcServer.SetDNSResponse, error) {
	return nil, nil
}
func (s *ServerAPI) DeleteDNS(context.Context, *grpcServer.DeleteDNSRequest) (*grpcServer.DeleteDNSResponse, error) {
	return nil, nil
}
