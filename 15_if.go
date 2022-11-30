package main

import "fmt"

func main() {
	name := "Bagaskara"

	if name == "Ramadhan" {
		fmt.Println("Where is Bagaskara?")
	} else if name == "Bagas" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Halo salam kenal.")
	}

	if length := len(name); length > 5{
		println("Nama terlalu panjang");
	}
}
