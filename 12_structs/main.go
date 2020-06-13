package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Initialize person using struct
	person1 := Person{firstName: "Rashid", lastName: "Razak", city: "Kuala Lumpur", gender: "Male", age: 29}
	fmt.Println(person1)

	// Alternative
	person2 := Person{"Shairah", "Rahman", "Kuala Lumpur", "Female", 29}
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	// Using value receiver. Refer greet() method at the bottom
	fmt.Println(person1.greet())

	// Using pointer receiver. Refer hasBirthday() method at the bottom
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())

	// Using pointer receiver again. Refer getMarried() method at the bottom
	fmt.Println(person2.greet())
	person2.getMarried(person1)
	fmt.Println(person2.greet())
}

// Person struct declaration
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greet method - This is a value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method - This is a pointer receiver
// Pointer receiver does not return anything. It is just changing value
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method - This is another pointer receiver
func (p *Person) getMarried(spouse Person) {
	if p.gender == "Male" {return}
	p.lastName = spouse.lastName
}
