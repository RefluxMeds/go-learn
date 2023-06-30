package main

import "fmt"

func main() {
	xs := make([]string, 0, 50)

	xs = append(xs, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut")

	fmt.Println("The length of the slice is: ", len(xs))
	fmt.Println("The capacity of the slice is: ", cap(xs))

	for i := 0; i < 7; i++ {
		fmt.Printf("Index: %v \t \t Value: %v\n", i, xs[i])
	}
}
