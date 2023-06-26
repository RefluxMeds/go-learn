package main

import (
	"fmt"

	"github.com/RefluxMeds/golang-learn/udemy-course-examples/ninja-level-13/dogs"
)

type canine struct {
	name string
	age  int
}

type human struct {
	name string
	age  int
}

func main() {
	dog1 := canine{
		name: "Fido",
		age:  dogs.ToDogYears(10),
	}

	human1 := human{
		name: "Jerry",
		age:  dogs.ToHumanYears(56),
	}

	fmt.Println(dog1)

	fmt.Println(human1)
}
