// +build example

package main

import (
	"encoding/json"
	"fmt"
	"github.com/Chewy-Inc/lets-go/aws/alb"
)

func main() {
	response, _ := alb.LambdaResponse(200, "OK!")
	jsonString, _ := json.Marshal(response)
	fmt.Println(string(jsonString))
}
