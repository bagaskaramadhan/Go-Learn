package main

import "fmt"

func main() {

	var name1 string = "Bagas"
	var name2 string = "Bagas"
	resultName := name1 == name2

	fmt.Println(resultName)

	var age1 int = 25
	var age2 int = 21

	var resultAge bool= age1 != age2

	fmt.Println(resultAge)

	var number1 int =20;
	var number2 int =21;

	resultNumber := number1<number2;
	fmt.Println(resultNumber)
}