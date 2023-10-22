package main

import "fmt"

func main() {
	// Arrays
	// has to be fixed length
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr, fruitArr[1])

	// Declare and assign at same time
	fruits := [2]string{"Banana", "Guava"}
	fmt.Println(fruits)

	// slices
	fruitSlice := []string{"Banana", "Guava", "Grape", "Cherry"}
	fmt.Println(fruitSlice, len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
