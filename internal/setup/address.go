package setup

import (
	"math"
	"net"
	"os"
	"strconv"
)

func MustAddress() string {
	ipAddr := os.Getenv("ECHO_ADDRESS")
	if ipAddr == "" {
		panic("ECHO_ADDRESS must be set")
	}

	if net.ParseIP(ipAddr) == nil {
		panic("ECHO_ADDRESS must be valid ip-address")
	}

	portStr := os.Getenv("ECHO_PORT")
	if portStr == "" {
		panic("ECHO_PORT must be set")
	}

	intPort, err := strconv.Atoi(portStr)
	if err != nil {
		panic("ECHO_PORT must be integer")
	}

	if intPort < 0 || intPort > math.MaxUint16 {
		panic("ECHO_PORT must be in range [0, 65535]")
	}

	return ipAddr + ":" + portStr
}
