package main

import "fmt"

func main() {
	r := generate()

	for v := range r {
		fmt.Println(v)
	}
}

func generate() chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
