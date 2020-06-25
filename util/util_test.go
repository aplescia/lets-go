package util_test

import (
	"github.com/aplescia-chwy/lets-go/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestGetFinalElementOfPath(t *testing.T) {
	path := "some/path/hello"
	result := util.GetFinalElementOfPath(path)
	assert.Equal(t, "hello", result)
}

func TestParseTimeStringAsTimeOrNil(t *testing.T) {
	timeString := "1994-09-19T09:00:00.312Z"
	result, err := util.ParseTimeStringAsTimeOrNil(timeString, time.RFC3339Nano)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, 1994, result.Year())
	assert.Equal(t, "September", result.Month().String())
	assert.Equal(t, 19, result.Day())
	nanoTime := result.Nanosecond()/1000000
	assert.Equal(t, 312, nanoTime)
	timeString = "hey, bobby!"
	result, err = util.ParseTimeStringAsTimeOrNil(timeString, time.RFC3339Nano)
	if err == nil {
		t.Fatal(err)
	}
	assert.Nil(t, result)
	timeString = "1994-09-19T09:00:00Z"
	result, err = util.ParseTimeStringAsTimeOrNil(timeString, time.RFC3339Nano)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, result)
}

func TestInitLoggerWithLevel(t *testing.T) {
	err := os.Setenv("LOG_LEVEL", "info")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	logObject, _ := util.InitLoggerWithLevel(nil)
	assert.Equal(t, logrus.InfoLevel.String(), logObject.Level.String())
	err = os.Setenv("LOG_LEVEL", "bobby")
	logObject, err = util.InitLoggerWithLevel(nil)
	assert.NotNil(t, err)
	level, err := logrus.ParseLevel("error")
	logObject, _ = util.InitLoggerWithLevel(&level)
	assert.Equal(t, logrus.ErrorLevel.String(), logObject.Level.String())
}
