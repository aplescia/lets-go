package rest

import (
	"bytes"
	"encoding/json"
	"github.com/google/jsonapi"
	"github.com/google/uuid"
	"net/http"
	"reflect"
)

//SerializeAsJsonApiDocument serializes the given struct pointer, or slice of struct pointers, as
//a jsonapi document. Must be annotated with proper jsonapi annotations.
//see: https://github.com/google/jsonapi#jsonapi-tag-reference for info on how to annotate your
//structs.
func SerializeAsJsonApiDocument(someJsonApiStruct interface{}) (string, error) {
	switch reflect.TypeOf(someJsonApiStruct).Kind() {
	case reflect.Slice:
		x := reflect.ValueOf(someJsonApiStruct)
		var newSlice []interface{}
		for i := 0; i < x.Len(); i++ {
			if !x.Index(i).IsZero() {
				thing := x.Index(i).Elem()
				if thing.Kind() == reflect.Struct {
					newSlice = append(newSlice, thing.Addr().Interface())
				} else if !thing.CanAddr() && thing.IsValid() && !thing.IsNil() {
					if thing.Kind() == reflect.Ptr {
						newSlice = append(newSlice, thing.Interface())
					}
				}
			}
		}
		var buf bytes.Buffer
		err := jsonapi.MarshalPayload(&buf, newSlice)
		return buf.String(), err
	default:
		var buf bytes.Buffer
		err := jsonapi.MarshalPayload(&buf, someJsonApiStruct)
		return buf.String(), err
	}
}

//UnmarshalJsonApiDocument takes a JSON API document (json) as a byte slice and attempts to deserialize the data
//into a json api conformant struct pointer. Returns any errors.
func UnmarshalJsonApiDocument(documentJson []byte, inputStruct interface{}) error {
	return jsonapi.UnmarshalPayload(bytes.NewReader(documentJson), inputStruct)
}

//UnmarshalManyJsonApiDocument unmarshals a json api document that contains many payloads and returns a slice of structs that match
//the type of the inputStruct. therefore, the inputStruct passed must be the same type as the struct
//mapped to the JsonAPI payload.
func UnmarshalManyJsonApiDocument(documentJson []byte, inputStruct interface{}) ([]interface{}, error) {
	slice, err := jsonapi.UnmarshalManyPayload(bytes.NewReader(documentJson), reflect.TypeOf(inputStruct))
	return slice, err
}

//JsonApiErrorResponse takes a status code and an error object and serializes a JSON:API conformant error
//document.
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

//MarshalAsJsonString marshals a struct as a jsonString. Returns any errors.
func MarshalAsJsonString(someInput interface{}) (string, error) {
	result, err := json.Marshal(someInput)
	return string(result), err
}

//UnmarshalJsonString unmarshals a JSON String into a passed struct. You should be passing in a pointer.
//Returns any errors.
//For example:
//	var structToBePopulated MyStruct
//	rest.UnmarshalJsonString(jsonString, &structToBePopulated)
//or
//	var structToBePopulated *MyStruct
//	rest.UnmarshalJsonString(jsonString, structToBePopulated)
func UnmarshalJsonString(someString string, object interface{}) error {
	if object == nil {
		panic("Object is nil!")
	}
	err := json.Unmarshal([]byte(someString), &object)
	return err
}
