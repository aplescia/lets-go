//+build example

package main

import (
	"fmt"
	"github.com/aplescia-chwy/lets-go/db/redis"
)

var (
	clientPtr = redis.ClusterClient()
)

func main() {
	result := clientPtr.Ping()
	fmt.Print(result)
	clientPtr.Set("name", "Bob Ross", 0)
}
