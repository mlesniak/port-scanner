package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
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
var port = flag.String(
	"port",
	"",
	"a single port (80) or a single range (80-1024)")

// Set after command line parsing.
var timeout time.Duration
var portList []int

func main() {
	parseCommandLine()
	scanPorts()
}

func scanPorts() {
	for _, port := range portList {
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
	if *hostname == "" || *port == "" {
		flag.Usage()
	}
	// Special handling for ports to determine if a range is set.
	if strings.Contains(*port, "-") {
		ps := strings.Split(*port, "-")
		start, err := strconv.Atoi(ps[0])
		if err != nil {
			flag.Usage()
		}
		end, err := strconv.Atoi(ps[1])
		if err != nil {
			flag.Usage()
		}
		portList = make([]int, end - start + 1)
		for i := 0; i < len(portList); i++ {
			portList[i] = start + i
		}
	} else {
		// Single port.
		sp, err := strconv.Atoi(*port)
		if err != nil {
			flag.Usage()
		}
		portList = make([]int, 1)
		portList[0] = sp
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
	defer conn.Close()
	return true
}
