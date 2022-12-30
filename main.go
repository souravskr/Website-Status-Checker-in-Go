package main

import (
	"fmt"
	"net/http"
	"os"
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

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}
}

func checkWebsite(site string, c chan string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
		c <- "The site might be down!"
		os.Exit(1)
	}
	fmt.Println(site, "status is", res.Status)
	c <- "The website is up!"
}
