package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}
func main() {
	a := 1
	b := 2
	a, b = swap(a, b)
	fmt.Println(a, "\t", b)
}
