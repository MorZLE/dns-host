package service

import (
	"context"
	grpcServer "dns-host/gen/server"
	"dns-host/pkg"
	"dns-host/pkg/cerror"
	"log/slog"
)

func NewService(log *slog.Logger, dns *DNSWorker) IService {
	return &service{log: log, dns: dns}
}

type IService interface {
	SetHostname(ctx context.Context, newHost string) error
	GetHostname(ctx context.Context) (string, error)
	RestartHost(ctx context.Context) error

	GetAllDNS(ctx context.Context) ([]*grpcServer.Dns, error)
	AddDNS(ctx context.Context, nameServer, ip string) error
	DeleteDNS(ctx context.Context, nameServer, ip string) error
	RestartDNS(ctx context.Context) error
}

type service struct {
	log *slog.Logger
	dns *DNSWorker
}

func (s *service) SetHostname(ctx context.Context, newHost string) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}
	if !pkg.Domain(newHost).Valid() {
		return cerror.ErrBadHostname
	}

	return setHostname(ctx, newHost)
}

func (s *service) GetHostname(ctx context.Context) (string, error) {
	if ctx.Err() != nil {
		return "", cerror.ErrCancelled
	}
	hostname, err := getHostname(ctx)
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func (s *service) RestartHost(ctx context.Context) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}
	return restartHostnamed()

}

func (s *service) GetAllDNS(ctx context.Context) ([]*grpcServer.Dns, error) {
	if ctx.Err() != nil {
		return nil, cerror.ErrCancelled
	}

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

func (s *service) AddDNS(ctx context.Context, nameServer, ip string) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}

	if !pkg.Ip(ip).Valid() {
		return cerror.ErrBadIP
	}
	if !pkg.Domain(nameServer).Valid() {
		return cerror.ErrBadHostname
	}

	err := s.dns.addDNS(ctx, nameServer, ip)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteDNS(ctx context.Context, nameServer, ip string) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}

	if nameServer == "" && ip == "" {
		return cerror.ErrBadDNS
	}

	err := s.dns.deleteDNS(ctx, nameServer, ip)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) RestartDNS(ctx context.Context) error {
	if ctx.Err() != nil {
		return cerror.ErrCancelled
	}
	return s.dns.restartManagerDNS()

}
