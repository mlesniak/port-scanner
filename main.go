package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// Command line parameter
var hostname = flag.String(
	"hostname",
	"",
	"hostname of the target system")
var timeoutSeconds = flag.Float64(
	"timeout",
	1.0,
	"Timeout in seconds. Fractional values, e.g. 0.5 are allowed")
var timeout time.Duration

func main() {
	parseCommandLine()

	for port := 70; port < 90; port++ {
		p := scanPort("tcp", *hostname, port, timeout)
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
	if *hostname == "" {
		flag.Usage()
	}

	// Convert timeout from fractional seconds to time.Duration.
	var ms = int(*timeoutSeconds * 1000)
	timeout = time.Duration(ms) * time.Millisecond
}

func scanPort(tcpType, hostname string, port int, timeout time.Duration) bool {
	fmt.Println("Trying", hostname+":", port)
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(tcpType, address, timeout)
	if err != nil {
		return false
	}
	// Should we use a defer here?
	conn.Close()
	return true
}
