package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		wg.Add(1)
		go checkLink(link)
	}

	wg.Wait()
}

func checkLink(link string) {
	resp, err := http.Get(link)
	if resp.StatusCode > 299 {
		fmt.Println("The page", link, "reponse failed with status code:", resp.StatusCode)
	} else {
		fmt.Println(link, "status code is", resp.StatusCode)
	}

	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
}
