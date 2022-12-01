package main

import "fmt"

func getFullName() (string, string) {
	return "Bagaskara", "Ramadhan"
}

func main() {
	_, lastName := getFullName()
	fmt.Println( lastName)
}