package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	p := sort.StringsAreSorted(strs)
	fmt.Println("Strings sorted: ", p)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Ints sorted: ", s)
}
