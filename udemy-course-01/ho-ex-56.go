package main

import "fmt"

func main() {
	anon := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Karlo",
		friends: map[string]int{
			"John":   27,
			"Austin": 35,
			"Joseph": 15,
		},
		favDrinks: []string{"Pina colada", "Ice tea", "Beer"},
	}

	fmt.Printf("%#v\n", anon)

	for k, _ := range anon.friends {
		fmt.Println(anon.first, "- friends - ", k)

	}

	for _, v := range anon.favDrinks {
		fmt.Println(anon.first, "- drinks - ", v)
	}
}
