package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
