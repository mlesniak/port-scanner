package main

import (
	"os"
	"fmt"
	"net"
	"strconv"
	"time"
	"flag"
)

var hostname = flag.String("hostname", "localhost", "hostname")

func main() {
	flag.Usage = func() {
		fmt.Println("A simple port scanner in go.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	
	for port := 70; port < 90; port++ {
		p := scanPort("tcp", *hostname, port, time.Second*1)
		if p {
			fmt.Println(port, ":", p)
		}
	}
}

func scanPort(tcpType, hostname string, port int, timeout time.Duration) bool {
	fmt.Println("Trying", hostname + ":", port)
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(tcpType, address, timeout)
	if err != nil {
		return false
	}
	// Should we use a defer here?
	conn.Close()
	return true
}
