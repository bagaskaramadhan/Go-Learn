package main

import "fmt"

func person(fullName string, age int) {
	fmt.Println("Name:", fullName, "Age:", age)
}

func main() {
	person("Bagaskara Ramadhan", 25)

	fullName := "Bagaskara"
	person(fullName, 25)
}
