package kinesis

import (
	"github.com/Chewy-Inc/lets-go/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

func PutToKinesis(streamName string, streamRegion string, input string) (bool, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(streamRegion),
	})
	if err != nil {
		panic(err)
	}
	kc := kinesis.New(s)
	streamNameConv := aws.String(streamName)
	_, err = kc.PutRecord(&kinesis.PutRecordInput{
		Data:         []byte(input),
		StreamName:   streamNameConv,
		PartitionKey: aws.String("key1"),
	})
	if err != nil {
		log.Error("record was not put successfully", err)
		return false, err
	}
	return true, nil
}

func PutManyRecordsToKinesis(streamName string, streamRegion string, inputs []string) (bool, error) {
	s, err := session.NewSession(&aws.Config{
		Region: &streamRegion,
	})
	if err != nil {
		panic(err)
	}
	kc := kinesis.New(s)
	var input []*kinesis.PutRecordsRequestEntry
	for _, i := range inputs {
		var entry = &kinesis.PutRecordsRequestEntry{
			Data:         []byte(i),
			PartitionKey: aws.String("key1"),
		}
		input = append(input, entry)
	}
	_, err = kc.PutRecords(&kinesis.PutRecordsInput{
		Records:    input,
		StreamName: &streamName,
	})
	if err != nil {
		log.Println("records were not put successfully", err)
		return false, err
	}
	return true, nil
}
