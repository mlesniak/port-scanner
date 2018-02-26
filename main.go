// Main file starting execution.
package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// scanResult defines the result of the scan for a single port. It's open value is true if
// a connection to the port could be established.
type scanResult struct {
	port int
	open bool
}

func main() {
	parseCommandLine()
	servicesFuture := parseServiceList()
	results := scanPorts()
	services := <-servicesFuture
	printResults(results, services)
}

// Scan the list of ports in the global variable portList by trying to open a TCP connection 
// (hardcoded), timeing out after timeout. If a connection could be established, it is 
// directly closed.
func scanPorts() map[int]bool {
	ports := make(chan scanResult)
	sem := make(Semaphore, *parallel)

	for _, port := range portList {
		go func(ports chan scanResult, port int) {
			// Block if more than *parallel connections are started.
			sem.Acquire(1)
			defer sem.Release(1)
			isOpen := scanPort("tcp", *hostname, port, timeout)
			ports <- scanResult{port, isOpen}
		}(ports, port)
	}

	// Transform the channel into a map for easier consumption.
	results := make(map[int]bool)
	for i := 0; i < len(portList); i++ {
		res := <-ports
		results[res.port] = res.open
	}
	return results
}


// scanPort scans a single port on hostname, aborting after timeout.
func scanPort(tcpType, hostname string, port int, timeout time.Duration) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(tcpType, address, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// printResults displays the result of the scan in a format similar to nmap.
func printResults(results map[int]bool, services map[int]string) {
	format := "%-9v %-7v %v\n"
	fmt.Printf(format, "PORT", "STATUS", "SERVICE")
	for _, port := range portList {
		state := results[port]
		var status string
		if state {
			status = "open"
		} else {
			status = "closed"
		}
		var portProtocol = fmt.Sprintf("%v", port) + "/tcp"
		var service = services[port]
		fmt.Printf(format, portProtocol, status, service)
	}
}
