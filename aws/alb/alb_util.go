package alb

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func LambdaToALBResponse(statuscode int, body string) (events.ALBTargetGroupResponse, error) {
	var returnPayload events.ALBTargetGroupResponse
	returnPayload.Body = body
	returnPayload.StatusCode = statuscode
	returnPayload.StatusDescription = http.StatusText(statuscode)
	returnPayload.IsBase64Encoded = false
	returnPayload.Headers = make(map[string]string)
	returnPayload.Headers["Content-Type"] = "application/json"

	return returnPayload, nil
}

func NotFoundResponse() (events.ALBTargetGroupResponse, error) {
	return LambdaToALBResponse(404, "")
}

func InternalServerError(err error) (events.ALBTargetGroupResponse, error) {
	return LambdaToALBResponse(500, err.Error())
}