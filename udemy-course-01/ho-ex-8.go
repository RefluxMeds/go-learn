package main

import "fmt"

func main() {
	var signed int8 = 127
	var unsigned uint8 = 255

	fmt.Printf("Value of signed: %d\n", signed)
	fmt.Printf("Type of signed: %T\n", signed)
	fmt.Printf("Value of unsigned: %d\n", unsigned)
	fmt.Printf("Type of signed: %T\n", unsigned)
}
