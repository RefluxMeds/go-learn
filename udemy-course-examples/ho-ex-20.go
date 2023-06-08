package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs.")
}

func main() {
	x := rand.Intn(500)
	fmt.Printf("The value of x is %v\n", x)

	switch {
	case x < 100:
		fmt.Println("between 0 and 100")
	case x > 100 && x < 200:
		fmt.Println("between 101 and 200")
	case x > 200 && x < 250:
		fmt.Println("between 200 and 250")
	default:
		fmt.Println("bigger than 250")
	}

}
