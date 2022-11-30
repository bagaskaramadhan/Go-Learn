package main

import "fmt"

func main() {
	var names[2]string;
	names[0] = "Bagaskara"
	names[1] = "Ramadhan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names)

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(len(names))
}