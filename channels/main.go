package main

import (
	"time"
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)
	for _, link := range links {
	 go checkLink(link, c)
	}
	
	for l :=  range c{ 
		//  This for loop says watch for channel c 
		//  and whenever a value comes from c assign it to l
		 
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
		// A function literal which is something like an 
		// anonymous function in java
	}
	// fmt.Println(<- c) // receiving a messgage from channel is blocking statement for the main routine
}



func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	fmt.Println("hiii") 
	if err != nil {
		c <- link 
	}
	c <- link
	fmt.Println(link, " is up")
}
