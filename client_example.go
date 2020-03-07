package main

import (
	"fmt"
	"github.com/SabaunT/shmedis/shmedis_sevice"
)

func main() {
	// port at which server is
	address := "8080"
	a := shmedis_sevice.Client(address)

	// "SET"
	a.Set("1", 123)

	// "GET"
	k := a.Get("1")
	fmt.Println("client got value", k.DataValue) // 123

	// "KEYS"
	fmt.Println("keys", a.Keys()) // [1]

	// "REMOVE"
	a.RemoveKey("1")
	fmt.Println("keys", a.Keys()) // []

	b := a.Get("1")
	fmt.Println("got", b.DataValue) // nil

	a.Close() // Connection to memecache server is closed.
}
