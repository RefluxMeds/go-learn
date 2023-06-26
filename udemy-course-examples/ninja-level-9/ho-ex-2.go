package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("My name is", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	h1 := person{
		first: "James",
	}

	h2 := person{
		first: "Bond",
	}

	saySomething(&h1)
	saySomething(&h2)
}
