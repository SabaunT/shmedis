package main

import (
	"fmt"
	"local/shmedis/memcache"
	"time"
)

func main() {
	duration, _ := time.ParseDuration("3s")
	memcache.NewCache(duration, duration)
	fmt.Scanln()
}
