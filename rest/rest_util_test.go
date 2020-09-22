package rest_test

import (
	"errors"
	"testing"

	"github.com/aplescia-chwy/lets-go/rest"
	"github.com/stretchr/testify/assert"
)

type MyTestStruct struct {
	ID   string `jsonapi:"primary,MyTestStruct"`
	Name string `jsonapi:"attr,name"`
}

type MyBadStruct struct {
	Name string
	Code int
}

func TestSerializeAsJsonApiResponse(t *testing.T) {
	testObj := &MyTestStruct{
		ID:   "10",
		Name: "bob",
	}
	output, err := rest.SerializeAsJSONAPIDocument(testObj)
	t.Log(output, err)
	output, err = rest.SerializeAsJSONAPIDocument(&MyBadStruct{
		Name: "",
		Code: 0,
	})
	t.Log(output, err)
	output, err = rest.SerializeAsJSONAPIDocument("hello")
	assert.NotNil(t, err)
	t.Log(output, err)
}

func TestUnmarshalJsonApiDocument(t *testing.T) {
	testObj := &MyTestStruct{
		ID:   "10",
		Name: "bob",
	}
	output, _ := rest.SerializeAsJSONAPIDocument(testObj)
	toBeDeserialized := new(MyTestStruct)
	err := rest.UnmarshalJSONAPIDocument([]byte(output), toBeDeserialized)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(toBeDeserialized)
	assert.Equal(t, toBeDeserialized, testObj)
}

func TestUnmarshalManyJsonApiDocument(t *testing.T) {
	testObj := &MyTestStruct{
		ID:   "10",
		Name: "bob",
	}
	testObjTwo := &MyTestStruct{
		ID:   "11",
		Name: "bobby",
	}
	var objs []*MyTestStruct
	objs = append(objs, testObj)
	objs = append(objs, testObjTwo)
	objs = append(objs, nil)
	output, err := rest.SerializeAsJSONAPIDocument(objs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(output)
	res, err := rest.UnmarshalManyJSONAPIDocument([]byte(output), testObjTwo)
	t.Log(len(res))
	resString, err := rest.SerializeAsJSONAPIDocument(res)
	t.Log(resString, err)
}

func TestJsonApiErrorResponse(t *testing.T) {
	err := errors.New("This is my test Error")
	res := rest.JSONAPIErrorResponse(500, err)
	t.Log(res)
	assert.NotEmpty(t, res)
}

func TestMarshalling(t *testing.T) {
	testObj := &MyTestStruct{
		ID:   "10",
		Name: "bob",
	}
	output, err := rest.MarshalAsJSONString(testObj)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(output)
	var thingy MyTestStruct
	err = rest.UnmarshalJSONString(output, &thingy)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(thingy)
}
