package main

import "fmt"

func greeting(name string) string {
	return "Hello, " + name + "!"
}

// both are int then we can directly write it as
// num1, num2 int
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Aman"))
	fmt.Println(getSum(4, 8))
}
