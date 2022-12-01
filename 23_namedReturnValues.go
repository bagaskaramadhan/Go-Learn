package main

import "fmt"

func getCompletedName() (age int, firstName, middleName, lastName string) {
	age = 26
	firstName = "Ibrahim"
	middleName = "Anakara"
	lastName = "Rendjana"
	return
}

func main() {
	age, name1, name2, name3 := getCompletedName()
	fmt.Println(age, name1, name2, name3)
}
