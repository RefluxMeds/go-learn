package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	n := add(5, 5)
	if n != 10 {
		log.Fatalf("Testing Add function failed. Got: %d. Wanted: %v", n, 10)
	}
}

func TestSubtract(t *testing.T) {
	n := subtract(5, 5)
	if n != 0 {
		log.Fatalf("Testing Subtract function failed. Got: %d. Wanted: %v", n, 0)
	}
}

func TestDoMath(t *testing.T) {
	n := doMath(1, 5, add)
	if n != 10 {
		log.Fatalf("Testing doMath function failed. Got: %d. Wanted: %v", n, 10)
	}

	m := doMath(5, 5, subtract)
	if m != 0 {
		log.Fatalf("Testing doMath function failed. Got: %d. Wanted: %v", m, 0)
	}
}
