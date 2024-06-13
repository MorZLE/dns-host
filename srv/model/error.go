package model

import "errors"

var (
	ErrBadIP       = errors.New("bad format ip")
	ErrBadHostname = errors.New("bad format hostname")
	ErrBadDNS      = errors.New("bad format dns")
)
