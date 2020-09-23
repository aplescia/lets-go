package sqs

import (
	"github.com/aplescia-chwy/lets-go/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	logger, _ = util.InitLoggerWithLevel(nil)
)

//BuildClient builds an SQS client using a passed AWS Config pointer.
func BuildClient(config *aws.Config) *sqs.SQS {
	lSession := session.Must(session.NewSession())

	lSvc := sqs.New(lSession, config)

	return lSvc
}

//PushToSqsAndReturnErrors pushes an input message to an SQS queue using a passed client object. Returns any errors.
func PushToSqsAndReturnErrors(client *sqs.SQS, input *sqs.SendMessageInput) error {
	_, err := client.SendMessage(input)
	if err == nil {
		logger.Printf("SQS Message Publish Successful To Queue %v! ", input.QueueUrl)
	}
	return err
}

func GetSqsEventLength(event events.SQSEvent) int {
	return len(event.Records)
}

func ProcessSqsEvent(event events.SQSEvent, messageProcessingFunc func(events.SQSMessage) error) {
	for _, r := range event.Records {
		e := messageProcessingFunc(r)
		if e != nil {
			logger.Errorf("%e", e)
		}
	}
}
