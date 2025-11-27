package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	fmt.Println(sumup(1, 10, 15, 40, -5))
	anotherSum := sumup(30, numbers...)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		fmt.Println(val)
		sum += val
	}
	fmt.Println(startingValue)

	return sum
}
