# shmedis - redis on minimals
Actually it's a for fun implementation of memcache service. Don't take this seriously.

## Features:
1. TCP connection between client and server;
2. "GET", "SET", "REMOVE", "CLOSE", "KEYS" operators;
3. TTL for keys;

## Use
server_example.go
```go
package main

import (
    "github.com/SabaunT/shmedis/shmedis_sevice"
    "time"
)

func main() {
    // port 
    address := "8080"

    // each 3 seconds server will run cache cleaner 
    cleanerInterval, _ := time.ParseDuration("3s")

    // TTL for each key
    expirationDuration, _ := time.ParseDuration("3s")
    shmedis_sevice.UpServer(address, cleanerInterval, expirationDuration)
}
```

client_example.go
```go
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
```

