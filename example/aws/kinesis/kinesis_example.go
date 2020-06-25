// +build example

package main

import "github.com/aplescia-chwy/lets-go/aws/kinesis"

func main() {
	someKinesisString := "Hey, buddy! This is my Kinesis Payload!"
	kinesis.PutToKinesis("my-stream-name", someKinesisString)
}
