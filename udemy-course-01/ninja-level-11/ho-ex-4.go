package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		seerr := sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  errors.New(fmt.Sprint("Hello from sqrtError struct err field.")),
		}
		return f, seerr
	}
	return 42, nil
}
