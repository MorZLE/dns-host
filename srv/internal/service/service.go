package service

import (
	"context"
	grpcServer "dns-host/gen/server"
	"dns-host/srv/model"
	"dns-host/srv/model/cerror"
	"log/slog"
	"strings"
)

func NewService(log *slog.Logger, dns *Dns) IService {
	return &service{log: log, dns: dns}
}

type IService interface {
	SetHostname(ctx context.Context, newHost string) error
	GetHostname(ctx context.Context) (string, error)
	GetAllDNS(ctx context.Context) ([]*grpcServer.Dns, error)
	SetDNS(ctx context.Context, name, ip string) error
	DeleteDNS(ctx context.Context, name, ip string) error
}

type service struct {
	log *slog.Logger
	dns *Dns
}

func (s *service) SetHostname(ctx context.Context, newHost string) error {
	if !model.Domain(newHost).Valid() {
		return cerror.ErrBadHostname
	}

	return setHostname(ctx, newHost)
}

func (s *service) GetHostname(ctx context.Context) (string, error) {
	hostname, err := getHostname(ctx)
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func (s *service) GetAllDNS(ctx context.Context) ([]*grpcServer.Dns, error) {
	var resp []*grpcServer.Dns

	mapDns, err := s.dns.getAllDNS(ctx)
	if err != nil {
		return nil, err
	}
	for k, v := range mapDns {
		resp = append(resp, &grpcServer.Dns{
			Ip:         string(k),
			NameServer: string(v),
		})
	}

	return resp, nil
}
func (s *service) SetDNS(ctx context.Context, name, ip string) error {
	if !model.Ip(ip).Valid() {
		return cerror.ErrBadIP
	}
	if name == "" {
		return cerror.ErrBadHostname
	}
	if len(strings.Fields(name)) > 1 {
		return cerror.ErrBadHostname
	}
	err := s.dns.setDNS(ctx, name, ip)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) DeleteDNS(ctx context.Context, name, ip string) error {
	if name == "" && ip == "" {
		return cerror.ErrBadDNS
	}
	err := s.dns.deleteDNS(ctx, name, ip)
	if err != nil {
		return err
	}

	return nil
}
