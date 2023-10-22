package main

import (
	"fmt"
	"strconv"
)

// importing string convertor
// define a struct
// No classes in golang
type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

// alternative
// type Person struct {
// 	firstname, lastname, city, gender string
// 	age                               int
// }

// two kind of methods
// value receivers and pointer receivers

// Greeting method (value receiver)
// we dont change anything here
func (p Person) greet() string {
	// similar to this keyword
	return "Hello!, My name is " + p.firstname + " " + p.lastname + ", and I am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method (Pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastname = spouseLastname
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstname: "Aman", lastname: "Varshney", city: "Etah", gender: "Male", age: 20}
	// Alternative
	person2 := Person{"Gauri", "Varshney", "Etah", "Female", 19}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstname)
	// can change the properties
	person1.age++
	fmt.Println(person1)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1)

	fmt.Println(person2)
	person2.getMarried("Gupta")
	fmt.Println(person2)
}
