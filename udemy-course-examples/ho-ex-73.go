package main

import (
	"fmt"
	"time"
)

func main() {
	TimedFunction(doWork)
}
func doWork() {
	for i := 0; i < 2000; i++ {
		fmt.Println(i)
	}
}

func TimedFunction(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
