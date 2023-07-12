package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":    "#ff0000",
		"blue":   "0000ff#",
		"green":  "#008000",
		"purple": "#800080",
	}

	fmt.Println(colors)

	boje := make(map[string]string)
	boje["white"] = "#FFFFFF"
	fmt.Println(boje)

	delete(colors, "green")

	printMap(colors)

}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("Hex code for %v is %v\n", k, v)
	}
}
