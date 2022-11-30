package main

import "fmt"

func main() {
	name := "Bagaskara"

	switch name {
	case "Ana":
		fmt.Println("Halo Ana")
	case "Ramadhan":
		fmt.Println("Halo Ramadhan")
	default:
		fmt.Println("Halo", name)
	}


	// switch length := len(name); length > 9 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah sesuai.")
	// }

	// short switch
	length := len(name);
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang.");
	case length > 5:
		fmt.Println("Nama lumayan panjang.")
	default:
		fmt.Println("Nama cukup panjang")
	}
}