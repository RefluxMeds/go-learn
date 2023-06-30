package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Your architecture is:", runtime.GOARCH)
	fmt.Println("Your operating system is:", runtime.GOOS)
}
