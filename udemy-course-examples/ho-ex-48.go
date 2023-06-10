package main

import "fmt"

func main() {
	jbxs := []string{"James", "Bond", "Shaken, no stirred"}
	mmxs := []string{"Miss", "Moneypenny", "I'm 008."}

	xxs := [][]string{jbxs, mmxs}

	fmt.Println(xxs)

	for i, _ := range xxs {
		for i, v := range xxs[i] {
			fmt.Printf("Index: %v \t\t Value: %v\n", i, v)
		}
	}
}
