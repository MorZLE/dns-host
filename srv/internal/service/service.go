package service

import (
	"context"
	grpcServer "dns-host/gen/server"
	"dns-host/srv/model"
	"log/slog"
	"net"
)

func NewService(log *slog.Logger) IService {
	return &service{log: log}
}

type IService interface {
	SetHostname(ctx context.Context, newHost string) error
	GetHostname(ctx context.Context) string
	GetAllDNS(ctx context.Context) []*grpcServer.Dns
	SetDNS(ctx context.Context, name, shortname, ip string) error
	DeleteDNS(ctx context.Context, name, ip string) error
}

type service struct {
	log *slog.Logger
}

func (s *service) SetHostname(ctx context.Context, newHost string) error {
	if newHost == "" {
		return model.ErrBadHostname
	}

	return nil
}

func (s *service) GetHostname(ctx context.Context) string {

	return "test"
}

func (s *service) GetAllDNS(ctx context.Context) []*grpcServer.Dns {
	return nil
}
func (s *service) SetDNS(ctx context.Context, name, shortname, ip string) error {
	if ipVal := net.ParseIP(ip).To4(); ipVal == nil {
		return model.ErrBadIP
	}
	if name == "" {
		return model.ErrBadHostname
	}

	return nil
}
func (s *service) DeleteDNS(ctx context.Context, name, ip string) error {
	if name == "" && ip == "" {
		return model.ErrBadDNS
	}

	return nil
}
