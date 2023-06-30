package main

import "fmt"

func returnFunc() func() int {
	return func() int {
		return 42
	}
}

func main() {
	x := returnFunc()
	fmt.Println(x())
}
