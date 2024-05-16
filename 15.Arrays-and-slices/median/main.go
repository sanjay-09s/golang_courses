package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 2, 4, 8, 1}
	fmt.Println(n)
	fmt.Println("Median-one : ", median_one(n))
	fmt.Println(n)
	fmt.Println("-------------")

	x := []float64{3, 2, 1, 4}
	fmt.Println(x)
	fmt.Println("Median-two : ", median_two(n))
	fmt.Println(x)
	fmt.Println("-------------")
}

func median_one(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i]
	}
	return (x[i-1] + x[i]) / 2
}

func median_two(x []float64) float64 {
	n := make([]float64, 6)
	copy(n, x)
	sort.Float64s(n)
	i := len(x) / 2

	if len(x)%2 == 1 {
		return x[i]
	}
	return (x[i-1] + x[i]) / 2
}
