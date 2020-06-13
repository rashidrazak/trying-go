package main

// Package name must be in double quotes
import "fmt"

// Shorthand will not work outside a function
var year int = 2020

func main() {
	// Note: All declared variables inside a function must be used

	// Declaring variables
	var firstname string = "Rashid"
	const lastname string = "Razak"
	var age = 29

	// Shorthand
	favouriteBand := "Westlife"

	fmt.Println("02_variables")
	fmt.Println(firstname, lastname, age)
	fmt.Println("My favourite band is", favouriteBand)

	// Show variable's type
	fmt.Printf("%T\n", favouriteBand)
}