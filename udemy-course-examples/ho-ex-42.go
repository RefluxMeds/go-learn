package main

import "fmt"

func main() {
	n := [5]int{1, 2, 3, 4, 5}

	for i, v := range n {
		fmt.Printf("Index: %v\t -> \tvalue: %v\n", i, v)
		fmt.Printf("Index type %T, value type %T\n", i, v)
	}
	fmt.Printf("Variable n is of %T type. Variable n holds value: %v.\n", n, n)

	m := [5]int{}
	for i := 0; i < 5; i++ {
		m[i] = i
	}
	for i, v := range m {
		fmt.Printf("Index: %v\t -> \tvalue: %v\n", i, v)
		fmt.Printf("Index type %T, value type %T\n", i, v)
	}

}
