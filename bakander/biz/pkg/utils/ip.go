package utils

import (
	"errors"
	"net"
)

// GetLocalIPv4Address 单网卡
func GetLocalIPv4Address() (string, error) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, add := range addr {
		if ipNet, ok := add.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", errors.New("not found ipv4 address")
}

// GetLocalIPv4AddressArr 多网卡模式
func GetLocalIPv4AddressArr() (ipAdds []string, err error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ipAdds, errors.New("not found ipv4 address")
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addr, _ := netInterfaces[i].Addrs()
			for _, address := range addr {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ipAdds = append(ipAdds, ipNet.IP.String())
					}
				}
			}
		}
	}
	return ipAdds, errors.New("not found ipv4 address")
}
