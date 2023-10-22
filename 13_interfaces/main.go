package main

import (
	"fmt"
	"math"
)

// datatypes represents set of method structures
// method signatures for structs

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	length, breadth float64
}

// area for circle (value receiver)
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// area for circle (value receiver)
func (r Rectangle) area() float64 {
	return r.breadth * r.length
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{length: 10, breadth: 7}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
