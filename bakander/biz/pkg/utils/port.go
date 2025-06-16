package utils

import (
	"net"
)

// GetFreePort get a free port.
func GetFreePort(TCP string, FreePortAddress string) (int, error) {
	addr, err := net.ResolveTCPAddr(TCP, FreePortAddress)
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP(TCP, addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
