package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(s string) {
	p.firstName = s
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Andersen",
	}

	alex.print()

	var james person
	james.firstName = "James"
	james.lastName = "Bond"

	jim := person{
		firstName: "Jim",
		lastName:  "Morisson",
		contactInfo: contactInfo{
			email:   "j.morisson@outlook.com",
			zipCode: 30000,
		},
	}

	jim.print()

	jim.updateName("Jimmy")
	jim.print()
}
