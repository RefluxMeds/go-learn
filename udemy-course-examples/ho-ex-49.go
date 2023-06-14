package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":       []string{"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": []string{"intelligence", "literature", "computer science"},
		"no_dr":            []string{"cats", "ice cream", "sunsets"},
	}

	for k, _ := range m {
		fmt.Println("For key", k)
		for k, v := range m[k] {
			fmt.Printf("Index position %v \t Value is %v\n", k, v)
		}
	}
}
