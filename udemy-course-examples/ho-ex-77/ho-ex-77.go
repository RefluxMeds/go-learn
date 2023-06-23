package main

import "fmt"

type person struct {
	first string
}

func value(p person, s string) person {
	p.first = s
	return p
}

func pointer(p *person, s string) {
	p.first = s
}

func main() {
	p1 := person{
		first: "Adrian",
	}
	fmt.Println(p1)

	p1 = value(p1, "Max")
	fmt.Println(p1)

	pointer(&p1, "Bruno")
	fmt.Println(p1)

	p2 := person{
		first: "Carmilla",
	}
	fmt.Println(p2)

	p2 = value(p2, "Betty")
	fmt.Println(p2)

	pointer(&p2, "Martina")
	fmt.Println(p2)
}
