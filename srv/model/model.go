package model

import "net"

type Domain string
type Ip string

func (domain Domain) Valid() bool {
	return domain != "" && len([]byte(domain)) <= 255
}

func (ip Ip) Valid() bool {
	return net.ParseIP(string(ip)).To4() != nil
}
