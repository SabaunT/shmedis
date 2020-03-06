package shmedis_sevice

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/SabaunT/shmedis/memcache"
	"github.com/SabaunT/shmedis/utils"
	"net"
	"time"
)

func UpServer(port string, cleanUpInterval, dataExpireAfter time.Duration) {
	address := fmt.Sprintf(":%v", port)
	listener, err := net.Listen("tcp", address)
	utils.HandleError(err)

	cache := memcache.NewCache(cleanUpInterval, dataExpireAfter)

	for {
		connection, err := listener.Accept()
		utils.HandleError(err)
		go handleConnection(connection, cache)
	}
}

func handleConnection(conn net.Conn, cache *memcache.Cache) {

	connectionScanner := bufio.NewScanner(conn)
	connectionWriter := json.NewEncoder(conn)

	for connectionScanner.Scan() {
		req := &utils.Request{}
		scannedMessage := connectionScanner.Bytes()
		err := json.Unmarshal(scannedMessage, req)
		utils.HandleError(err)

		if req.Method == "SET" {
			fmt.Println("Got SET command with args:", req.Arguments)
			cache.Set(req.Arguments.Key, req.Arguments.Value)
		}

		if req.Method == "GET" {
			fmt.Println("Got GET. Returning value under key", req.Arguments.Key)
			ret := cache.Get(req.Arguments.Key)

			err := connectionWriter.Encode(ret)
			utils.HandleError(err)
		}

		if req.Method == "KEYS" {
			fmt.Println("Got KEYS requests")
			ret := cache.Keys()

			err := connectionWriter.Encode(ret)
			utils.HandleError(err)
		}

		if req.Method == "REMOVE" {
			fmt.Println("Got REMOVE for key", req.Arguments.Key)
			cache.RemoveKey(req.Arguments.Key)
		}

		if req.Method == "CLOSE" {
			fmt.Println("Deleting cache")
			memcache.DeleteCache(cache)
		}
	}
}
