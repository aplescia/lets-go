package alb

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

//LambdaToALBResponse creates an AWS ALB compatible response payload dependent on a given
//status code and body message. Returns any errors.
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

//NotFoundResponse creates an AWS ALB compatible Not Found payload with an empty body. Returns any errors.
func NotFoundResponse() (events.ALBTargetGroupResponse, error) {
	return LambdaToALBResponse(404, "")
}

//InternalServerError creates an AWS ALB compatible Internal Server Error payload given a Go error object. Returns any errors.
func InternalServerError(err error) (events.ALBTargetGroupResponse, error) {
	return LambdaToALBResponse(500, err.Error())
}
