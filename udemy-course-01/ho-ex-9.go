package main

import "fmt"

var i int = 5

const j = "hello"

func main() {
	k := true

	fmt.Printf("%v is type %T\n%v is type %T\n%v is type %T\n", i, i, j, j, k, k)
}
