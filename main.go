package main

import (
	"os"
	"fmt"
	"net"
	"strconv"
	"time"
	"flag"
)

var hostname = flag.String("hostname", "", "hostname of the target system")

func main() {
	parseCommandLine()
	
	for port := 70; port < 90; port++ {
		p := scanPort("tcp", *hostname, port, time.Second*1)
		if p {
			fmt.Println(port, ":", p)
		}
	}
}

func parseCommandLine() {
	flag.Usage = func() {
		fmt.Println("A simple port scanner in go.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	// Check that mandatory arguments are defined.
	if (*hostname == "") {
		flag.Usage()
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
