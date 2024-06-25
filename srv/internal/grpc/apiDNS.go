package grpc

import (
	"context"
	grpcServer "dns-host/gen/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *ServerAPI) AddDNS(ctx context.Context, req *grpcServer.AddDNSRequest) (*grpcServer.AddDNSResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	grpclog.Info("add dns", slog.String("hostname", req.NameServer), slog.String("ip", req.Ip))
	var resp grpcServer.AddDNSResponse

	err := s.srv.AddDNS(ctx, req.NameServer, req.Ip)
	if err != nil {
		grpclog.Error("failed to add dns", err)
		return &resp, err
	}
	return &resp, nil
}

func (s *ServerAPI) GetAllDNS(ctx context.Context, req *grpcServer.GetAllDNSRequest) (*grpcServer.GetAllDNSResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	grpclog.Info("get all dns")
	var resp grpcServer.GetAllDNSResponse

	items, err := s.srv.GetAllDNS(ctx)
	if err != nil {
		grpclog.Error("failed to get all dns", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp.Items = items

	return &resp, nil
}

func (s *ServerAPI) DeleteDNS(ctx context.Context, req *grpcServer.DeleteDNSRequest) (*grpcServer.DeleteDNSResponse, error) {
	if ctx.Err() != nil {
		return nil, status.Error(codes.Canceled, "the client canceled the request")
	}

	grpclog.Info("delete dns", slog.String("hostname", req.NameServer))
	var resp grpcServer.DeleteDNSResponse

	err := s.srv.DeleteDNS(ctx, req.NameServer, req.Ip)
	if err != nil {
		grpclog.Error("failed to delete dns", err)
		return &resp, err
	}

	return &resp, nil
}
