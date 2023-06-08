package main

import "fmt"

type ByteSize int

func main() {
	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	fmt.Printf("%d \t\t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t\t %b\n", TB, TB)
	fmt.Printf("%d \t\t %b\n", PB, PB)
	fmt.Printf("%d \t\t %b\n", EB, EB)
}
