package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("The area of %T is: %v\n", s, s.getArea())
}

func main() {
	trokut := triangle{
		height: 2.3,
		base:   5.1,
	}

	kocka := square{
		sideLength: 3.4,
	}

	printArea(trokut)
	printArea(kocka)
}
