package cmd

import (
	"dns-host/client/internal"
)

var service internal.ClientDNSHost

func InitGRPCinCLI(srv internal.ClientDNSHost) {
	service = srv
}
