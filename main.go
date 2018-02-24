package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	p631 := scanPort("tcp", "localhost", 631, time.Second*5)
	fmt.Println("631:", p631)
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
