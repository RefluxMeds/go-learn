package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)
	fmt.Printf("Value of x is %v\nValue of y is %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater or equal to 4 and less or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}
}
