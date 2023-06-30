package main

import "fmt"

var zeroval string

const specific = 100

func main() {
	sdo := 10
	i, j, z := 1, true, "hello"
	_ = "nista"

	fmt.Printf("Zeroval is of type %T and value %v.\n", zeroval, zeroval)
	fmt.Printf("Constant is of type %T and value %v.\n", specific, specific)
	fmt.Printf("Short assignment: %v\n Multiple init i: %T, %v\n Multiple init j: %T, %v\n Multiple init z: %T, %v\n", sdo, i, i, j, j, z, z)
}
