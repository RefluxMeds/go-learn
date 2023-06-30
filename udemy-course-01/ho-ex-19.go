package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v\n", x)

	if x < 100 {
		fmt.Println("between 0 and 100")
	} else if x <= 200 && x >= 101 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 201 and 250")
	}

	fmt.Println(rand.Intn(3))
}
