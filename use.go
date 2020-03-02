package main

import (
	"fmt"
	"time"
	"local/medis/memcache"
)

func main() {
	duration, _ := time.ParseDuration("3s")
	memcache.NewCache(duration, duration)
	fmt.Scanln()
}
