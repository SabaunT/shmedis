package main

import (
	"fmt"
	"local/shmedis/shmedis_client"
)

func main() {
	address := "8080"
	a := shmedis_client.Client(address)
	a.Set("1", 123)
	fmt.Println("1")
	k := a.Get("1")
	fmt.Println("client got value", k.DataValue)
	fmt.Println("keys", a.Keys())
	a.RemoveKey("1")
	fmt.Println("keys", a.Keys())
	b := a.Get("1")
	fmt.Println("got", b.DataValue)
	a.Set("1", 1234)
	a.Close()
	fmt.Println("keys", a.Keys())
}
