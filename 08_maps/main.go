package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)

	// assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// delete from Bob
	delete(emails, "Bob")
	fmt.Println(emails)

	// assign while declaring
	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(emails2)
}
