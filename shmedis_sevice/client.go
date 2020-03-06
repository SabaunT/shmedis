package shmedis_sevice

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/SabaunT/shmedis/memcache"
	"github.com/SabaunT/shmedis/utils"
	"net"
)

type Shmedis struct {
	conn        net.Conn
	connScanner *bufio.Scanner
	connEncoder *json.Encoder
}

func Client(port string) *Shmedis {
	address := fmt.Sprintf(":%v", port)
	connection, err := net.Dial("tcp", address)
	utils.HandleError(err)

	shmedisClient := &Shmedis{
		conn:        connection,
		connScanner: bufio.NewScanner(connection),
		connEncoder: json.NewEncoder(connection),
	}
	return shmedisClient
}

func (shmedisClient *Shmedis) Get(key string) *memcache.Data {
	arguments := utils.Arguments{
		Key:   key,
		Value: nil,
	}
	request := utils.Request{
		Method:    "GET",
		Arguments: arguments,
	}
	err := shmedisClient.connEncoder.Encode(request)
	utils.HandleError(err)

	ret := &memcache.Data{}
	for shmedisClient.connScanner.Scan() {
		scannedMessage := shmedisClient.connScanner.Bytes()
		err := json.Unmarshal(scannedMessage, ret)
		utils.HandleError(err)
		break
	}
	return ret
}

func (shmedisClient *Shmedis) Set(key string, value interface{}) {
	arguments := utils.Arguments{
		Key:   key,
		Value: value,
	}
	request := utils.Request{
		Method:    "SET",
		Arguments: arguments,
	}
	err := shmedisClient.connEncoder.Encode(request)
	utils.HandleError(err)
}

func (shmedisClient *Shmedis) Keys() []string {
	request := utils.Request{
		Method:    "KEYS",
		Arguments: utils.Arguments{},
	}
	err := shmedisClient.connEncoder.Encode(request)
	utils.HandleError(err)

	ret := new([]string)
	for shmedisClient.connScanner.Scan() {
		scannedMessage := shmedisClient.connScanner.Bytes()
		err := json.Unmarshal(scannedMessage, ret)
		utils.HandleError(err)
		break
	}
	return *ret
}

func (shmedisClient *Shmedis) RemoveKey(key string) {
	arguments := utils.Arguments{
		Key:   key,
		Value: nil,
	}
	request := utils.Request{
		Method:    "REMOVE",
		Arguments: arguments,
	}
	err := shmedisClient.connEncoder.Encode(request)
	utils.HandleError(err)
}

func (shmedisClient *Shmedis) Close() {
	request := utils.Request{
		Method:    "CLOSE",
		Arguments: utils.Arguments{},
	}
	err := shmedisClient.connEncoder.Encode(request)
	utils.HandleError(err)

	shmedisClient.conn.Close()
	fmt.Println("Connection to memecache server is closed.")
}
