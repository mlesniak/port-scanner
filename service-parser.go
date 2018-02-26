// Parse the embedded service list to provide a map from port to the port's service description.
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
)

// serviceFuture is a future (a channel with a single write) containing the map from 
// port number to abbreviated description of the port.
type serviceFuture chan map[int]string

func parseServiceList() chan map[int]string {
	future := make(serviceFuture)
	go internalParse(future)
	return future
}

func internalParse(future serviceFuture) {
	services := make(map[int]string)

	// Read data.
	bs, err := Asset("data/service-names-port-numbers.csv")
	if err != nil {
		fmt.Println(err)
		panic("Unable to access service file list")
	}

	// Parse CSV.
	r := csv.NewReader(bytes.NewReader(bs))
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create map.
	for _, rec := range records {
		// Remove UDP information for now.
		if rec[2] != "tcp" {
			continue
		}
		// Remove empty port numbers
		if rec[1] == "" {
			continue
		}
		port, err := strconv.Atoi(rec[1])
		if err != nil {
			// We ignore port ranges such as xwindow's 6000-6003 for now.
		}
		techDesc := rec[0]
		services[port] = techDesc
	}

	future <- services
}
