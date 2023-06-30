package main

import (
	"fmt"
)

func main() {
	goNum := 10
	c := make(chan int)

	for i := 0; i < goNum; i++ {
		go func() {
			for i := 0; i < goNum; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
}
