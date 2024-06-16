package grpc

import (
	"context"
	grpcServer "dns-host/gen/server"
	"dns-host/srv/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *ServerAPI) SetHostname(ctx context.Context, req *grpcServer.SetHostnameRequest) (*grpcServer.SetHostnameResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	s.log.Info("set hostname", slog.String("hostname", req.Hostname))
	err := s.srv.SetHostname(ctx, req.Hostname)
	if err != nil {
		s.log.Error("failed to set hostname", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &grpcServer.SetHostnameResponse{Success: true}, nil
}

func (s *ServerAPI) GetHostname(ctx context.Context, req *grpcServer.GetHostnameRequest) (*grpcServer.GetHostnameResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	s.log.Info("get hostname")
	var resp grpcServer.GetHostnameResponse
	hostname, err := s.srv.GetHostname(ctx)
	if err != nil {
		s.log.Error("failed to get hostname", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp.Hostname = hostname

	return &resp, nil
}

func (s *ServerAPI) AddDNS(ctx context.Context, req *grpcServer.AddDNSRequest) (*grpcServer.AddDNSResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	s.log.Info("add dns", slog.String("hostname", req.NameServer), slog.String("ip", req.Ip))
	var resp grpcServer.AddDNSResponse

	err := s.srv.AddDNS(ctx, req.NameServer, req.Ip)
	if err != nil {
		resp.Error = err.Error()
		s.log.Error("failed to add dns", err)
		return &resp, err
	}
	resp.Success = true

	return &resp, nil
}

func (s *ServerAPI) GetAllDNS(ctx context.Context, req *grpcServer.GetAllDNSRequest) (*grpcServer.GetAllDNSResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	s.log.Info("get all dns")
	var resp grpcServer.GetAllDNSResponse

	items, err := s.srv.GetAllDNS(ctx)
	if err != nil {
		s.log.Error("failed to get all dns", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp.Items = items

	return &resp, nil
}

func (s *ServerAPI) DeleteDNS(ctx context.Context, req *grpcServer.DeleteDNSRequest) (*grpcServer.DeleteDNSResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	s.log.Info("delete dns", slog.String("hostname", req.NameServer))
	var resp grpcServer.DeleteDNSResponse

	err := s.srv.DeleteDNS(ctx, req.NameServer, req.Ip)
	if err != nil {
		s.log.Error("failed to delete dns", err)
		resp.Error = err.Error()
		return &resp, err
	}
	resp.Success = true

	return &resp, nil
}
