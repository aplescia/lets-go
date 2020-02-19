package util_test

import (
	"github.com/Chewy-Inc/lets-go/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetFinalElementOfPath(t *testing.T) {
	path := "some/path/hello"
	result := util.GetFinalElementOfPath(path)
	assert.Equal(t, "hello", result)
}

func TestParseTimeStringAsTimeOrNil(t *testing.T) {
	timeString := "1994-09-19T09:00:00.312Z"
	result := util.ParseTimeStringAsTimeOrNil(&timeString)
	assert.NotNil(t, result)
	assert.Equal(t, 1994, result.Year())
	assert.Equal(t, "September", result.Month().String())
	assert.Equal(t, 19, result.Day())
	assert.Equal(t, 312, result.Nanosecond()/1_000_000)
	timeString = "hey, bobby!"
	result = util.ParseTimeStringAsTimeOrNil(&timeString)
	assert.Nil(t, result)
	timeString = "1994-09-19T09:00:00Z"
	result = util.ParseTimeStringAsTimeOrNil(&timeString)
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
