package main

import (
	"fmt"
	"math"
)

// func (r receiver) identifier(p parameters(s)) (r return(s)) { code }

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	Area() float64
}

func (s square) Area() float64 {
	n := s.length * s.width
	return n
}

func (c circle) Area() float64 {
	n := math.Pi * math.Pow(c.radius, 2)
	return n
}

func info(s shape) float64 {
	return s.Area()
}

func main() {
	square := square{
		length: 5.5,
		width:  3.2,
	}

	circle := circle{
		radius: 3,
	}

	fmt.Println(info(square))

	fmt.Println(info(circle))
}
