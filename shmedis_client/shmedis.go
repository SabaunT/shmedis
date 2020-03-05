package shmedis_client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"local/shmedis/memcache"
	"local/shmedis/util"
	"net"
)

type Shmedis struct {
	connScanner *bufio.Scanner
	encoder *json.Encoder
	decoder *json.Decoder
}

func Client(port string) *Shmedis {
	address := fmt.Sprintf(":%v", port)
	connection, err := net.Dial("tcp", address)
	util.HandleError(err)

	shmedisClient := &Shmedis{
		connScanner: bufio.NewScanner(connection),
		encoder: json.NewEncoder(connection),
	}
	return shmedisClient
}

func (shmedisClient *Shmedis) Get(key string) *memcache.Data {
	arguments := util.Arguments{
		Key: key,
		Value: nil,
	}
	request := util.Request{
		Method: "GET",
		Arguments: arguments,
	}
	err := shmedisClient.encoder.Encode(request)
	util.HandleError(err)

	ret := &memcache.Data{}
	for shmedisClient.connScanner.Scan() {
		scannedMessage := shmedisClient.connScanner.Bytes()
		err := json.Unmarshal(scannedMessage, ret)
		util.HandleError(err)
		break
	}
	return ret
}

func (shmedisClient *Shmedis) Set(key string, value interface{}) {
	arguments := util.Arguments{
		Key: key,
		Value: value,
	}
	request := util.Request{
		Method: "SET",
		Arguments: arguments,
	}
	err := shmedisClient.encoder.Encode(request)
	util.HandleError(err)
}
//
//func (shmedisClient *Shmedis) Keys() []string {
//	return shmedisClient.Keys()
//}
//
//func (shmedisClient *Shmedis) RemoveKey(key string) {
//	shmedisClient.RemoveKey(key)
//}
//
//func (shmedisClient *Shmedis) Close() {
//	memcache.DeleteCache(shmedisClient.Cache)
//	shmedisClient.serviceListener.Close()
//}
