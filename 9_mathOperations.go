package main

import "fmt"

func main() {
	var a = 10;
	var b = 10;
	var c = a + b;

	fmt.Println(c)

	// Augmented
	a += 5 // a + 5
	fmt.Println(a)

	// Unary
	a++ // a + 1
	fmt.Println(a)
}