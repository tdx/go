package api

import (
	"io"
	"net"
)

// Resolver is an interface for host name ipv4 resolver
type Resolver interface {
	AddHost(host string)
	DelHost(host string)
	GetNextIP(host string) string
	GetIPs(host string) ([]net.IP, []net.IP)
	GetIPsStr(host string) ([]string, []string)
	Dump(w io.Writer)
}
