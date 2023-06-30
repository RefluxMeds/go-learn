package main

import "fmt"

func main() {
	y := 0
	for {
		fmt.Printf("y has value of %v\n", y)
		if y > 20 {
			break
		}
		y++
	}
}
