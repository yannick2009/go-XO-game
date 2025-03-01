package utils

import "net"

// GetMyIP returns the IP address of the current machine in the local network
// It returns an empty string if the IP address is not found
func GetMyIP() string {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, ip := range ips {
		if ipnet, ok := ip.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
