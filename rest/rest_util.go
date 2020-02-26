package rest

import (
	"bytes"
	"github.com/Chewy-Inc/lets-go/util"
	"github.com/google/jsonapi"
	"net/http"
	"github.com/google/uuid"

)

var (
	log, _ = util.InitLoggerWithLevel(nil)
)

func SerializeAsJsonApiResponse(someJsonApiStruct *interface{}) string {
	var buf bytes.Buffer
	err := jsonapi.MarshalPayload(&buf, someJsonApiStruct)
	if err != nil {
		log.Panicln(err)
	}
	return buf.String()
}

func JsonApiErrorResponse(statusCode int, err error) string {
	randomUuid, _ := uuid.NewRandom()
	httpStatus := http.StatusText(statusCode)
	obj := &jsonapi.ErrorObject{
		ID:     randomUuid.String(),
		Title:  "Error",
		Detail: err.Error(),
		Status: httpStatus,
		Code:   "4001",
		Meta:   nil,
	}
	var buf bytes.Buffer
	var errs []*jsonapi.ErrorObject
	errs = append(errs, obj)
	_ = jsonapi.MarshalErrors(&buf, errs)
	return buf.String()
}