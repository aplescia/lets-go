package concurrency

import (
	"github.com/aplescia-chwy/lets-go/util"
	"sync"
)

var (
	logger, _ = util.InitLoggerWithLevel(nil)
)

//ProcessEachElementOfSliceInParallel will run processingFunc on each element of slice inputSlice
//using goroutines. Logs each error at ERROR level.
func ProcessEachElementOfSliceInParallel(inputSlice []interface{}, processingFunc func(interface{}) error){
	wg := &sync.WaitGroup{}
	for _, i := range inputSlice {
		wg.Add(1)
		go func(i interface{}) {
			defer wg.Done()
			err := processingFunc(i)
			if err != nil {
				logger.Errorf("%e", err)
			}
		}(i)
		wg.Wait()
	}
}


