package iputils

import "net"

type Container interface {
	Contains(ip net.IP) bool
}

func Contains(slice []Container, ip net.IP) (ok bool) {
	for _, haser := range slice {
		if haser.Contains(ip) {
			return true
		}
	}
	return
}
