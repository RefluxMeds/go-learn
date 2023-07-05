package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range xi {
		if v%2 == 0 {
			fmt.Printf("Number %v is even\n", v)
		} else {
			fmt.Printf("Number %v is odd\n", v)
		}
	}
}
