package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 6)

	userNames = append(userNames, "Ryan")
	userNames = append(userNames, "Dwiky")
	userNames[0] = "Maudina"
	userNames[1] = "Khaerani"

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	for idx, value := range userNames {
		fmt.Println(idx)
		fmt.Println(value)
	}

	for key, value := range courseRatings {
		fmt.Println(key)
		fmt.Println(value)
	}
}
