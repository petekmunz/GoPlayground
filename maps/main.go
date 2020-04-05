package main

import "fmt"

func main() {
	//Different ways of initializing a map below

	//var colors map[string]string
	//colors := make(map[string]string)
	//colors["green"] = "#008000"

	//We can delete map values by use of the delete function
	//delete(colors, "green")

	colors := map[string]string{
		"blue":  "#0000FF", //Unlike other languages, in Go a comma is required after every key-value entry
		"red":   "#FF0000",
		"green": "#008000",
	}

	printMap(colors)
}

func printMap(cMap map[string]string) {
	//How to iterarte over a map below
	for key /*color*/, value /*hexcode*/ := range cMap {
		fmt.Println("Hexcode for ", key, " is ", value)
	}
}
