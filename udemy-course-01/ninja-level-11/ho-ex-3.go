package main

import "fmt"

type customErr struct {
	message string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", c.message)
}

func main() {
	ce := customErr{
		message: "This is a custom error struct.",
	}

	foo(ce)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")
	fmt.Println("assertion:", e.(customErr).message)
}
