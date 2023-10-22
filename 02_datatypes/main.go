package main

import (
	"fmt"
)

// global variable
var language = "Golang"

func main() {
	fmt.Println("Hello, World!")
	// MAIN TYPES
	// string, bool, int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte -- alias for uint8
	// rune -- alias for int32
	// float32, float64
	// complex64 complext128

	// using var keyword
	// specfying a datatype is not necessary
	var name string = "Aman"
	var surname = "Varshney"

	// shorthand method (this type of declaration must be inside a function)
	firstname := "Aman"
	height := 1.57 // by default float64
	// we can specifcally define float32
	var size float32 = 5.7

	var age = 19
	var age2 int32 = 19
	fmt.Println(name, firstname, surname, age, size)
	fmt.Printf("%T\n%T\n%T\n%T\n%T\n", size, height, age, age2, firstname)

	var isCool = false
	// const isCool = true  can not reassign the const variable just like JS
	isCool = true
	fmt.Println(isCool, language)
	fmt.Printf("%T\n", isCool)

	// other shorthands
	email, username := "aman@gmail.com", "amanVar06"
	fmt.Println(email, username)
}
