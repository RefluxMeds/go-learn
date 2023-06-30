package main

import "fmt"

func main() {

	defer fmt.Println("First print")
	defer fmt.Println("Second print")
	defer fmt.Println("Third print")
	defer fmt.Println("Fourth print")
	defer fmt.Println("Fifth print")

	fmt.Println("First no defer")
	fmt.Println("Second no defer")
}
