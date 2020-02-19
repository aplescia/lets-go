package util

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

//Get an environment variable value by key or return some default value.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//Return the final element of a path of the root.
//Example path: /this/is/my/path returns path.
func GetFinalElementOfPath(inputPath string) string {
	var pathMembers = strings.Split(inputPath, "/")
	return pathMembers[len(pathMembers)-1]
}

//Parse a time string in RFC3339Nano format as either a go time.Time object or nil.
//Example format: "2009-01-02T01:02:32.111Z"
func ParseTimeStringAsTimeOrNil(timeString *string) *time.Time {
	if timeString == nil || *timeString == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *timeString)
	if err != nil {
		log.Error("Could not parse time")
		return nil
	}
	return &t
}

//Creates a pointer to a logrus object. The logger defaults to the DEBUG level unless another logging
//level is passed as a function argument, or is specified under the environment variable LOG_LEVEL.
//the logger is configured to log to stdout. Returns an error if any.
func InitLoggerWithLevel(level *log.Level) (*log.Logger,error) {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	if level == nil {
		loggingLevelFromEnv := GetEnv("LOG_LEVEL", log.DebugLevel.String())
		level, err := log.ParseLevel(loggingLevelFromEnv)
		if err != nil {
			log.Error("Couldn't instantiate logger with custom level! Returning logger because : ", err)
			return logger, err
		} else {
			logger.SetLevel(level)
		}
	} else {
		logger.SetLevel(*level)
	}
	return logger, nil
}
