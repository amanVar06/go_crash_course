package main

import "fmt"

func main() {
	a := 5
	b := &a

	val := "aman"
	ptr := &val
	fmt.Println(a, b, ptr)
	fmt.Printf("%T\n%T\n", b, ptr)

	// Use * to read val from address
	fmt.Println(*b, *ptr)
	// *b === *&a === a

	// change val with pointer
	*b = 10
	fmt.Println(a)
}
