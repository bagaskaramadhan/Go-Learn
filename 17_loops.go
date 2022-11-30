package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// simple loop
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Bagas", "Ana", "Anakara"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// BEFORE USE i as [KEY] THEN USE _ FOR INITIATE DECLARE BUT NOT USED
	for _, value := range slice {
		fmt.Println("Ini adalah", value)
	}

	person := make(map[string]string)
	person["name"]="Bagaskara Ramadhan"
	person["title"]="Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
