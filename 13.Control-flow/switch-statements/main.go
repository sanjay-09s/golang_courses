package main

import "fmt"

func main() {
	x := 31
	y := 25

	// for one method
	switch {
	case x == 31:
		fmt.Println("x is equal")
	case x > 31:
		fmt.Println("x is greater")
	case x < 31:
		fmt.Println("x is smaller")
	default:
		fmt.Println("Default run there is ! error !")
	}

	// method 2
	switch y {
	case 12:
		fmt.Println("y is 12")
	case 21:
		fmt.Println("y is 21")
	case 29:
		fmt.Println("y is 29")
	default:
		fmt.Println("y is ", y)
	}
}
