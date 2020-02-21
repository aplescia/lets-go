// +build example

package main

import (
	"fmt"
	"github.com/Chewy-Inc/lets-go/db/redis"
)

var (
	client = redis.ClusterClient()
)

func main() {
	result := client.Ping()
	fmt.Print(result)
	client.Set("name","Bob Ross", 0)
}
