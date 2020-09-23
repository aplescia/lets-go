package sqs_test

import (
	"fmt"
	"os"
	"testing"

	sq "github.com/aplescia-chwy/lets-go/aws/sqs"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	testQueueUrl   = os.Getenv("TEST_QUEUE_URL")
	testAwsAccount = os.Getenv("TEST_AWS_ACCOUNT")
)

func TestPushToSqsAndReturnErrors(t *testing.T) {
	client := sq.BuildClient(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", testAwsAccount),
	})

	t.Log(client)
	err := sq.PushToSqsAndReturnErrors(client, &sqs.SendMessageInput{
		MessageBody: aws.String("Hello"),
		QueueUrl:    aws.String(testQueueUrl),
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestProcessSqsEvent(t *testing.T) {
	msg := events.SQSMessage{
		Body: "hey",
	}
	var records []events.SQSMessage
	records = append(records, msg)
	event := events.SQSEvent{
		Records: records,
	}
	sq.ProcessSqsEvent(event, processingFunc)
}

func processingFunc(event events.SQSMessage) error {
	fmt.Println(event.Body)
	return nil
}
