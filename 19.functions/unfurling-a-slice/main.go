package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := sum(x...)
	fmt.Println(sum)
}

func sum(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}
