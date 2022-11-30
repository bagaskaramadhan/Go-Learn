package main

import "fmt"

func main() {
	type idCard string
	type married bool

	var noKTP idCard = "3174091401971001"
	var marriedStatus married = true

	fmt.Println(noKTP)
	fmt.Println(marriedStatus)
}