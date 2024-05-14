package main

import "fmt"

func main() {

	var arr = []int{1, 2, 3, 4, 5, 6, 5, 64}
	for i, j := range arr {
		fmt.Printf("Ranging over the slice %v : %v\n", i, j)
	}

	// Maps
	s := map[string]int{
		"Sugan":  18,
		"Sanjay": 21,
	}

	for i, j := range s {
		fmt.Printf("Ranging over the map %v : %v\n", i, j)
	}

}
