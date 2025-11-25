package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age

	fmt.Println("Age:", age)
	// fmt.Println("Adult years:", getAdultYears(age))
	fmt.Println("Age without pointer:", age)

	fmt.Println("Age pointer:", agePointer)
	getAdultYears(agePointer)
	fmt.Println("Age with pointer:", age)
	fmt.Println("Age with pointer:", *agePointer)
}

func getAdultYears(age *int) {
	*age -= 18
}
