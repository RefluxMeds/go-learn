package main

import "fmt"

type person struct {
	first     string
	last      string
	ice_cream []string
}

func main() {
	p1 := person{
		first:     "Karlo",
		last:      "Bond",
		ice_cream: []string{"chocolate", "stracatella", "vanilla"},
	}

	p2 := person{
		first:     "Luka",
		last:      "Blajvić",
		ice_cream: []string{"caramel", "strawberry", "lemon"},
	}

	fmt.Printf("Variable p1 has type %T \t\t value %v\n", p1, p1)
	fmt.Printf("Variable p2 has type %T \t\t value %v\n", p2, p2)

	for _, v := range p1.ice_cream {
		fmt.Println(p1.first, "favorite is", v)
	}

	for _, v := range p2.ice_cream {
		fmt.Println(p2.first, "favorite is", v)
	}
}
