package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1]) // [10.99]
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices) // [10.99 9.99 5.99]
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	fmt.Println(prices)
// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices[3])

// 	featuredPrices := prices[1:3] // print [9.99 45.99]
// 	fmt.Println(featuredPrices)
// 	featuredPrices = prices[:3]
// 	fmt.Println(featuredPrices) // print [10.99 9.99 45.99]
// 	featuredPrices = prices[1:]
// 	fmt.Println(featuredPrices) // print [9.99 45.99 20]
// 	featuredPrices = prices[:4]
// 	fmt.Println(featuredPrices) // print [10.99 9.99 45.99 20]

// 	// chaining slices
// 	featuredPrices = prices[1:] // [9.99 45.99 20]
// 	fmt.Println(featuredPrices)
// 	highlightedPrices := featuredPrices[:1] // [9.99]
// 	fmt.Println(highlightedPrices)
// }
