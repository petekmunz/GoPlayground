package main

import "fmt"

type bot interface {
	//Interfaces are a contract to help us manage types
	//We can thus reuse code through interfaces as seen with both english & swahili bot in printGreeting()

	getGreeeting() string //For a type to satisfactorily be of type bot, it needs to fulfill the function declared in the interface
}

type englishBot struct{}
type swahiliBot struct{}

func main() {
	eBot := englishBot{}
	sBot := swahiliBot{}

	printGreeting(eBot)
	printGreeting(sBot)
}

func (eBot englishBot) getGreeeting() string { //Since englishBot can call getGreeting, it satisfies the func required to be type bot
	return "Hello!"
}

func (sBot swahiliBot) getGreeeting() string { //Since swahiliBot can call getGreeting, it also satisfies the func required to be type bot
	return "Jambo!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeeting())
}
