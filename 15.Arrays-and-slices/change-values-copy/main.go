package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a
	fmt.Println("-------------")
	fmt.Println(a)
	fmt.Println(b)

	a[0] = 7
	fmt.Println("-------------")
	fmt.Println(a)
	fmt.Println(b)

	c := make([]int, 6)
	copy(c, a)

	a[0] = 0

	fmt.Println("-------------")
	fmt.Println(a)
	fmt.Println(c)

}
