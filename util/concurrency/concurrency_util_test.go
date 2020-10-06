package concurrency_test

import (
	"fmt"
	"github.com/aplescia-chwy/lets-go/util/concurrency"
	"testing"
)

func TestProcessEachElementOfSliceInParallel(t *testing.T) {
	var strings = []interface{}{"hey", "there"}
	concurrency.ProcessEachElementOfSliceInParallel(strings, processingFunc)
}

func processingFunc(input interface{}) error {
	fmt.Println(input)
	return nil
}
