package main

import "fmt"

var a = "a"

func main() {

	fmt.Println(a)

	var x string = "one" // cannot be single quotes
	fmt.Println(x)

	var name = "ashu" // infers the type automatically
	fmt.Println(name)

	var will_use string // can define without assigning value
	fmt.Println(will_use)

	my_full_name := "Ashutosh Mahesh Pednekar" // shorthand to create and assign value
	// This cannot be used outside a function
	fmt.Println(my_full_name)

	// Strings can be reassigned, but the type cannot change
	// Similarly for other data-types...

}
