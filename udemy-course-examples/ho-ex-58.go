package main

import "fmt"

// func (r receiver) identifier(p parameters(s)) (r return(s)) { code }

func main() {

	fmt.Println(foo(5))
	fmt.Println(bar(5, "Anthony"))

	x := foo(10)
	fmt.Println(x)
	y, z := bar(50, "Marc")
	fmt.Println(y, z)

}

func foo(x int) int {
	return x
}

func bar(x int, name string) (int, string) {
	return x, name
}
