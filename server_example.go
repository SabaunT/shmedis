package main

import (
	"github.com/SabaunT/shmedis/shmedis_sevice"
	"time"
)

func main() {
	address := "8080"
	duration, _ := time.ParseDuration("3s")
	shmedis_sevice.UpServer(address, duration, duration)
}
