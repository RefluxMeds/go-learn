package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors string
	color string
}

func main() {
	auto1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "VW",
		model: "Tiguan",
		doors: "Four doors",
		color: "Red",
	}

	auto2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Somemodel",
		doors: "Two doors",
		color: "Blue",
	}

	fmt.Println(auto1)
	fmt.Println(auto2)

	fmt.Println(auto1.model)
	fmt.Println(auto2.color)

	fmt.Println(auto1.electric)
	fmt.Println(auto2.engine)
}
