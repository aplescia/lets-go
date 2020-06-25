// +build example

package main

import (
	"encoding/json"
	"fmt"
	"github.com/aplescia-chwy/lets-go/aws/alb"
)

func main() {
	response, _ := alb.LambdaToALBResponse(200, "OK!")
	jsonString, _ := json.Marshal(response)
	fmt.Println(string(jsonString))
}
