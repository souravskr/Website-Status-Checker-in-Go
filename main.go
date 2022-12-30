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

	for _, website := range websites {
		checkWebsite(website)
	}
}

func checkWebsite(site string) {
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(site, "status is", res.Status)
}
