package main

import "fmt"

type contactInfo struct {
	email       string
	countryCode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo //We can embed one struct into another
}

func main() {
	//Three Different ways of declaring structs below.

	//	sue := person{ "Sue", "Mueni"}  //Pass in names on assumption of order of values in struct
	//	sue := person{firstname: "Sue", lastname: "Mueni"} //explicitly using the struct values to define a person

	var sue person //Declare a struct var, then access the values within it
	sue.firstname = "Sue"
	sue.lastname = "Mueni"
	sue.contact.email = "sue@exampleMail.com"
	sue.contact.countryCode = 254

	//Create a pointer before updating the value
	//	suePointer := &sue //'&' operator requires the memory address of the value this variable(sue) is pointing at.
	//	suePointer.updateName("Susan")

	//Alternatively Go handles the change to pointer by itself hence we can call same function with var of person
	sue.updateName("Susan")
	sue.print()
}

//Go being a pass by value language, pointers are required for updating value in memory.
func (pointerToPerson *person) updateName(newFirstName string) {
	// '*' before a type e.g *person is a type description, in our case a pointer that only points to a person type
	// '*' operator before a pointer(memory address) requires the actual value the memory address is pointing at
	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
