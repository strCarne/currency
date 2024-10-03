package setup

import (
	"math"
	"net"
	"os"
	"strconv"
)

func MustAddress() string {
	ipAddr := os.Getenv("ADDRESS")
	if ipAddr == "" {
		panic("ADDRESS must be set")
	}

	if net.ParseIP(ipAddr) == nil {
		panic("ADDRESS must be valid ip-address")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		panic("PORT must be set")
	}

	intPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic("PORT must be integer")
	}

	if intPort < 0 || intPort > math.MaxUint16 {
		panic("PORT must be in range [0, 65535]")
	}

	return ipAddr + ":" + portStr
}
