package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
	factRec := factorialRecursion(5)
	fmt.Println(factRec)
}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

func factorialRecursion(number int) int {
	if number == 0 {
		return 1
	}
	fmt.Println("Number: ", number)
	return number * factorialRecursion(number-1)
}
