package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Andersen",
	}
	fmt.Println(alex)

	var james person

	fmt.Println(james)
}
