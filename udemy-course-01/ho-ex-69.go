package main

import "fmt"

var x = func() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	x()

	y := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	y()
}
