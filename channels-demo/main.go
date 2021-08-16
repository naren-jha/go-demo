package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://twitter.com",
	}

	c := make(chan string) // create a channel to communicate value of string

	for _, link := range links {
		go checkLink(link, c) // launch a new go routine, with the channel c
	}

	// block main go routine, by making it wait for all responses to come through
	// the number of responses we are expecting through the channel is the number of links that we have
	// for i := 0; i < len(links); i++ {
	// fmt.Println(<-c) // blocking call - to receive message from channel
	// }

	for l := range c {
		// function literal (similar to anonomous function in JS or lambda in Java)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // invoke function literal with the value
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
