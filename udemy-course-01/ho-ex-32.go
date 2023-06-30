package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age := m["James"]
	fmt.Println(age)

	age = m["Q"]
	fmt.Println(age)

	if q, ok := m["Q"]; !ok {
		fmt.Println("There is no Q key. This is a zero value of int type:", q)
	}
}
