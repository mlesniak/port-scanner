package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	for port := 0; port < 1024; port++ {
		p := scanPort("tcp", "localhost", port, time.Second*5)
		if p {
			fmt.Println(port, ":", p)
		}
	}
}

func scanPort(tcpType, hostname string, port int, timeout time.Duration) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(tcpType, address, timeout)
	if err != nil {
		return false
	}
	// Should we use a defer here?
	conn.Close()
	return true
}
