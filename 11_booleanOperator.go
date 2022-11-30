package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusNilaiAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsensi

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusNilaiAbsensi)
	fmt.Println(lulus)
}