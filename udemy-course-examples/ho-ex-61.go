package main

import "fmt"

// func (r receiver) identifier(p parameters(s)) (r return(s)) { code }

type person struct {
	first string
	age   int
}

func (p person) speak(name string, age int) {
	fmt.Println("My name is", name, "and I am", age, "years old.")
}

func (p person) Speak() {
	fmt.Println("My name is", p.first, "and my age is", p.age)
}

func main() {
	p1 := person{
		first: "John",
		age:   25,
	}

	p2 := person{
		first: "James",
		age:   27,
	}

	p1.speak(p1.first, p1.age)
	p2.Speak()
}

/*
OR
*/
