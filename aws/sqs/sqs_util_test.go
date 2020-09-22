package sqs_test

import (
	"os"
	"testing"

	sq "github.com/aplescia-chwy/lets-go/aws/sqs"
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
