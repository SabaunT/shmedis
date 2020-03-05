package main

import (
	"fmt"
	"local/shmedis/shmedis_client"
)

func main() {
	address := "8080"
	a := shmedis_client.Client(address)
	a.Set("1", 123)
	k := a.Get("1")
	fmt.Println("client got value", k.DataValue)
}
