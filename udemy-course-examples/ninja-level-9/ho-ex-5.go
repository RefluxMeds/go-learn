package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&counter))
}
