package main

import "fmt"

func main() {
	// Long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// 	// i++
	// }

	// Short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	// fizzbuzz
	for n := 1; n <= 100; n++ {
		if n%15 == 0 {
			fmt.Println(n, "Fizzbuzz")
		} else if n%3 == 0 {
			fmt.Println(n, "Fizz")
		} else if n%5 == 0 {
			fmt.Println(n, "Buzz")
		} else {
			fmt.Println(n)
		}
	}
}
