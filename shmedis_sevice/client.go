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
	conn net.Conn
}

func Client(port string) *Shmedis {
	address := fmt.Sprintf(":%v", port)
	connection, err := net.Dial("tcp", address)
	utils.HandleError(err)

	shmedisClient := &Shmedis{
		conn: connection,
	}
	return shmedisClient
}

func (shmedisClient *Shmedis) Get(key string) *memcache.Data {
	request := shmedisClient.createRequestBody("GET", key, nil)
	shmedisClient.sendRequest(request)

	ret := &memcache.Data{}
	shmedisClient.encapsulateServerResponseIn(ret)
	return ret
}

func (shmedisClient *Shmedis) Set(key string, value interface{}) {
	request := shmedisClient.createRequestBody("SET", key, value)
	shmedisClient.sendRequest(request)
}

func (shmedisClient *Shmedis) Keys() []string {
	request := shmedisClient.createRequestBody("KEYS", "", nil)
	shmedisClient.sendRequest(request)

	ret := new([]string)
	shmedisClient.encapsulateServerResponseIn(ret)
	return *ret
}

func (shmedisClient *Shmedis) RemoveKey(key string) {
	request := shmedisClient.createRequestBody("REMOVE", key, nil)
	shmedisClient.sendRequest(request)
}

func (shmedisClient *Shmedis) Close() {
	request := shmedisClient.createRequestBody("CLOSE", "", nil)
	shmedisClient.sendRequest(request)

	shmedisClient.conn.Close()
	fmt.Println("Connection to memecache server is closed.")
}

func (shmedisClient *Shmedis) createRequestBody(method, key string, value interface{}) utils.Request {
	req := utils.Request{
		Method: method,
		Arguments: utils.Arguments{
			Key:   key,
			Value: value,
		},
	}
	return req
}

func (shmedisClient *Shmedis) sendRequest(request utils.Request) {
	connectionEncoder := json.NewEncoder(shmedisClient.conn)
	err := connectionEncoder.Encode(request)
	utils.HandleError(err)
}

func (shmedisClient *Shmedis) encapsulateServerResponseIn(in interface{}) {
	connectionScanner := bufio.NewScanner(shmedisClient.conn)
	for connectionScanner.Scan() {
		scannedMessage := connectionScanner.Bytes()
		err := json.Unmarshal(scannedMessage, in)
		utils.HandleError(err)
		break
	}
}
