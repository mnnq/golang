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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		//you can set go routines only in front of function calls
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		//getting messages from a channel is a blocking thing, it blocks the main routine
		fmt.Println(<-c)
	}
	/*  both are the same
	for {
		go checkLink(<-c, c)
	}

	for l := range c {
		//if we set time.Sleep this will block the main routine everytime for 5 seconds and we
		//dont want that, maybe a child routine finished but it will have to wait until the main
		//routine gets back from the sleep
		//time.Sleep( 5 * time.Second) //time.Second is a type, this will stop for 5 seconds
		go checkLink(l, c)
	}
	*/

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, " is up!")
	c <- "Is up GG"
}
