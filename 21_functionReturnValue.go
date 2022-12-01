package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "No name detected"
	} else {
		return "Halo " + name
	}
	// return "Hello " + name
	// fmt.Println()
}

func main() {
	result := sayHello("Bagaskara")
	fmt.Println(result)

	fmt.Println(sayHello(""))

}
