package main

import "fmt"

func main() {
	x := 5
	y := 10

	// practice is not to use parentheses
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "purple"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is neither red nor blue")
	}

	// Switch
	color = "red"
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is neither red nor blue")
	}

	// && and
	// || or
}
