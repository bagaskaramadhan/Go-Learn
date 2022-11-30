package main

import "fmt"

func main() {
	// ######################
	var name string
	name = "Bagaskara"
	fmt.Println(name)

	name = "Bagaskara Ramadhan"
	fmt.Println(name)
	// ######################

	var name2 string = "Ramadhan"
	fmt.Println(name2)

	name2 = "Bagaskara Ramadhan"
	fmt.Println(name2)
	// ######################

	var age int = 25
	fmt.Println(age)
	// ######################

	name3 := "Bagaskara"
	fmt.Println(name3)

	name3 = "Bagaskara Ramadhan"
	fmt.Println(name3)
	
	// ######################
	var (
		firstName = "Bagaskara"
		lastName = "Ramadhan"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
