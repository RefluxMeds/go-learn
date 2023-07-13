package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
