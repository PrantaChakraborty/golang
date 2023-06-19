package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string) // make a channel to communicate with goroutines

	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c) // print the channel value
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		//fmt.Println(link, "might be down")
		c <- "Might be down i think" // sending message to the channel
		return
	}
	//fmt.Println(link, " is ok")
	c <- "Yep, the link is ok" // sending message to the channel
}
