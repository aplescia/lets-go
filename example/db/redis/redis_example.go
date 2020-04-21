//+build example

package main

import (
	"fmt"
	"github.com/Chewy-Inc/lets-go/db/redis"
	"github.com/aws/aws-sdk-go/aws"
)

var (
	clientPtr = redis.ClusterClient(aws.String("dummypassword"))
)

func main() {
	result := clientPtr.Ping()
	fmt.Print(result)
	clientPtr.Set("name", "Bob Ross", 0)
}
