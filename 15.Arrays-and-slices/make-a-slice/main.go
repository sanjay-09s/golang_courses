package main

import "fmt"

func main() {
	// Normal for making a slice
	si := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(si)
	fmt.Println("--------------")

	// make a slice using the make function
	// make(type, length, capacity)
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, 10, 11, 12, 13, 14, 15)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	yi := make([]int, 5, 30)
	fmt.Println(yi)
	fmt.Println(len(yi))
	fmt.Println(cap(yi))

}
