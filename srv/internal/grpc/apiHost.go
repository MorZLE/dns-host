package grpc

import (
	"context"
	grpcServer "dns-host/gen/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *ServerAPI) SetHostname(ctx context.Context, req *grpcServer.SetHostnameRequest) (*grpcServer.SetHostnameResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	grpclog.Info("set hostname", slog.String("hostname", req.Hostname))
	err := s.srv.SetHostname(ctx, req.Hostname)
	if err != nil {
		grpclog.Error("failed to set hostname", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &grpcServer.SetHostnameResponse{Success: true}, nil
}

func (s *ServerAPI) GetHostname(ctx context.Context, req *grpcServer.GetHostnameRequest) (*grpcServer.GetHostnameResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	grpclog.Info("get hostname")
	var resp grpcServer.GetHostnameResponse
	hostname, err := s.srv.GetHostname(ctx)
	if err != nil {
		grpclog.Error("failed to get hostname", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp.Hostname = hostname

	return &resp, nil
}

func (s *ServerAPI) RestartHost(ctx context.Context, req *grpcServer.RestartHostRequest) (*grpcServer.RestartHostResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}
	grpclog.Info("restart host")
	err := s.srv.RestartHost(ctx)
	if err != nil {
		grpclog.Error("failed to restart host", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcServer.RestartHostResponse{Success: true}, nil
}
