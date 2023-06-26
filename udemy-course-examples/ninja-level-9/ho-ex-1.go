package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Hello from 1st goroutine")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second / 2)
		fmt.Println("Hello from 2nd goroutine")
		wg.Done()
	}()
	wg.Wait()
}
