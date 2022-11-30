package main

import "fmt"

func main() {
	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	slice:=months[7:11]
	fmt.Println(months)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// ############################
	
	slice2 := months[11:]
	fmt.Println(slice2)
	// ############################

	slice3 := append(slice2, "Bagas")
	fmt.Println(slice3)
}