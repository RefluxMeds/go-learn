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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("For key %v\n", k)
		for i, v := range v.ice_cream {
			fmt.Printf("Index position %v \t value %v\n", i, v)
		}
	}
}
