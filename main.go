package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Hello, Port Scan")
	conn, err := net.DialTimeout("tcp", "localhost:631", time.Second * 5)
	if err != nil {
		fmt.Println("Connection timeout:", err)
		return
	}
	fmt.Println("Closing connection.")
	conn.Close()
}