package shmedis_server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"local/shmedis/util"
	"local/shmedis/memcache"
	"net"
	"time"
)

func Up(port string, cleanUpInterval, dataExpireAfter time.Duration) {
	address := fmt.Sprintf(":%v", port)
	listener, err := net.Listen("tcp", address)
	util.HandleError(err)

	cache := memcache.NewCache(cleanUpInterval, dataExpireAfter)
	for {
		connection, err := listener.Accept()
		util.HandleError(err)
		go handleConnection(connection, cache)
	}
}

func handleConnection(conn net.Conn, cache *memcache.Cache) {

	connectionScanner := bufio.NewScanner(conn)
	connectionWriter := json.NewEncoder(conn)

	for connectionScanner.Scan() {
		req := &util.Request{}
		scannedMessage := connectionScanner.Bytes()
		err := json.Unmarshal(scannedMessage, req)
		util.HandleError(err)

		if req.Method == "SET" {
			fmt.Println("Got SET command with args:", req.Arguments)
			cache.Set(req.Arguments.Key, req.Arguments.Value)
		}

		if req.Method == "GET" {
			fmt.Println("Got GET. Returning value under key", req.Arguments.Key)
			ret := cache.Get(req.Arguments.Key)

			err := connectionWriter.Encode(ret)
			util.HandleError(err)
		}
	}
}
