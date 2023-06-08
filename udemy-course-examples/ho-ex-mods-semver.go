package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	puppy.From11()
	puppy.From12()
	puppy.From13()

	fmt.Println(s1, s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
