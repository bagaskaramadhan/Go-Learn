package main

import "fmt"

func main() {
	person := map[string]string{
		"name":"Bagaskara",
		"address":"Depok",
	}

	person["title"] = "Programmer"
	
	person["title"] = "President"
	fmt.Println(person)
	fmt.Println(len(person))

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Bagaskara Ramadhan"
	fmt.Println(book)
}