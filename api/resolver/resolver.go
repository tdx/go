package api

import (
	"io"
	"net"
)

// ResolverService is an interface for host name ipv4 resolver
type ResolverService interface {
	AddHost(host string)
	DelHost(host string)
	GetNextIP(host string) string
	GetIPs(host string) ([]net.IP, []net.IP)
	GetIPsStr(host string) ([]string, []string)
	Dump(w io.Writer)
}
