package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(s string) {
			time.Sleep(5 * time.Second)
			checkLink(s, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(time.Second)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode > 299 {
		fmt.Println("The page", link, "reponse failed with status code:", resp.StatusCode)
	} else {
		fmt.Println(link, "status code is", resp.StatusCode)
	}

	c <- link
}
