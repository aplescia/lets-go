package util

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

//Get an environment variable value by key or return some default value.
func GetEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//Return the final element of a path of the root. Makes use of the filepath library.
//Example path: /this/is/my/path returns path.
func GetFinalElementOfPath(inputPath string) string {
	return filepath.Base(inputPath)
}

//Parse a time string in the given format as either a go time.Time object or nil.
//Example format: "2009-01-02T01:02:32.111Z"
func ParseTimeStringAsTimeOrNil(timeString string, layout string) (*time.Time, error) {
	if timeString == "" || layout == "" {
		return nil, errors.New("time input, or layout, were empty strings")
	}
	t, err := time.Parse(layout, timeString)
	if err != nil {
		return nil, err
	}
	return &t, err
}

//Creates a pointer to a logrus object. The logger defaults to the DEBUG level unless another logging
//level is passed as a function argument, or is specified under the environment variable LOG_LEVEL.
//the logger is configured to log to stdout. Returns an error if any.
func InitLoggerWithLevel(level *log.Level) (*log.Logger, error) {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	if level == nil {
		loggingLevelFromEnv := GetEnvOrDefault("LOG_LEVEL", log.DebugLevel.String())
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
