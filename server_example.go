package main

import (
	"github.com/SabaunT/shmedis/shmedis_sevice"
	"time"
)

func main() {
	// port
	address := "8080"

	// each 3 seconds server will run cache cleaner
	cleanerInterval, _ := time.ParseDuration("3s")

	// TTL for each key
	expirationDuration, _ := time.ParseDuration("3s")
	shmedis_sevice.UpServer(address, cleanerInterval, expirationDuration)
}