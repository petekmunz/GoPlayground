package main

import (
	"fmt"
	"net/http"
	"time"
)

//Channels are used to communicate between goRoutines.

func main() {
	links := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://facebook.com",
		"https://golang.org",
		"https://twitter.com",
	}

	channel := make(chan string) //Create a channel & specify its type value.

	for _, link := range links {
		go checkLink(link, channel) //Create a new go routine for every link to check the status of the websites
	}

	for l := range channel { //Infinite loop that proceeds as long as we receive a value back from the channel.

		//Do not reference the same exact value between goRoutines hence pass 'link' as parameter in function Literal
		go func(link string) { //Function Literal(Anonymous function) that allows for sleeping the current goRoutine
			time.Sleep(5 * time.Second) //for 5 secs
			checkLink(link, channel)
		}(l)
	}

	//We receive a value from a channel as below
	// example_Value := <- channel
}

func checkLink(link string, channel chan string) { //Website status checker function
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link //Syntax for passing info to a channel
		return
	}

	if resp != nil {
		defer resp.Body.Close() //Must close to avoid mem-leak & reuse http connection for another request
	}

	fmt.Println(link, "is up!")
	channel <- link
}
