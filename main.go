package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, website := range websites {
		go checkWebsite(website, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkWebsite(link, c)
		}(l)
	}
}

func checkWebsite(site string, c chan string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
		c <- site
		os.Exit(1)
	}
	fmt.Println(site, "status is", res.Status)
	c <- site
}
