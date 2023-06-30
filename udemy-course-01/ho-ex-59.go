package main

// broj 153
import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(ii ...int) int {
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

func bar(ii []int) int {
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
