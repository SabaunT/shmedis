package main

import (
	"local/shmedis/shmedis_server"
	"time"
)

func main() {
	address := "8080"
	duration, _ := time.ParseDuration("3s")
	shmedis_server.Up(address, duration, duration)
}
