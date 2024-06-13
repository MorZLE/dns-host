package service

import (
	"log/slog"
)

func NewService(log *slog.Logger) IService {
	return &service{log: log}
}

type IService interface {
	SetHostname(newHost string) error
	GetHostname() string
	GetAllDNS() [][]string
	SetDNS(name, ip string) error
	DeleteDNS(name service) error
}

type service struct {
	log *slog.Logger
}

func (s *service) SetHostname(newHost string) error {
	return nil
}

func (s *service) GetHostname() string {
	return ""
}

func (s *service) GetAllDNS() [][]string {
	return nil
}
func (s *service) SetDNS(name, ip string) error {
	return nil
}
func (s *service) DeleteDNS(name service) error {
	return nil
}
