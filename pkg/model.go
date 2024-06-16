package pkg

import (
	"net"
	"strings"
)

type Domain string
type Ip string

func (domain Domain) Valid() bool {
	return domain != "" && len([]byte(domain)) <= 255 && len(strings.Fields(string(domain))) == 1
}

func (ip Ip) Valid() bool {
	return net.ParseIP(string(ip)) != nil
}
